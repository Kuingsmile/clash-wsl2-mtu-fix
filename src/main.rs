use pnet::datalink;
use std::process::Command;
use std::error::Error;

const IP_TO_FIND: &str = "198.18.0.1";

fn main() -> Result<(), Box<dyn Error>> {
    let interfaces = datalink::interfaces();

    let interface = interfaces.iter().find(|interface| {
        interface.ips.iter().any(|ip| ip.ip().to_string() == IP_TO_FIND)
    });

    match interface {
        Some(interface) => {
            println!("Found the interface: {}", interface.name);

            let output = Command::new("sudo")
                .arg("ip")
                .arg("link")
                .arg("set")
                .arg(&interface.name)
                .arg("mtu")
                .arg("1400")
                .output()?;

            let stdout = String::from_utf8_lossy(&output.stdout);
            let stderr = String::from_utf8_lossy(&output.stderr);
            println!("stdout: {}", stdout);
            println!("stderr: {}", stderr);
        },
        None => println!("No interface found with IP {}", IP_TO_FIND),
    }

    Ok(())
}