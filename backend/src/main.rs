use std::{net::{TcpStream, TcpListener}, io::prelude::*, fs::File};
#[macro_use]
extern crate log;

fn main() {
    env_logger::init();
    let listener = TcpListener::bind("0.0.0.0:9090").unwrap();

    for stream in listener.incoming() {
        let stream = stream.unwrap();

        return_contents(stream);
    }
}


fn return_contents(mut stream: TcpStream) {
    let mut request = [0; 1024];

    stream.read(&mut request).unwrap();
    debug!("Request: {}", String::from_utf8_lossy(&request[..]));

    let mut file = File::open("example.json").unwrap();
    let mut body = String::new();
    file.read_to_string(&mut body).unwrap();

    let response = format!("HTTP/1.1 200 OK\r\n\r\n{}", body);
    stream.write(response.as_bytes()).unwrap();
    stream.flush().unwrap();

}