import { program } from "commander";
import { green, red } from "kleur/colors";
import net from "net";

function performHealthCheck(domain: string, port: string): Promise<string> {
    const address = `${domain}:${port}`;
    const timeout = 5000;

    return new Promise((resolve) => {
        const socket = net.createConnection({ host: domain, port: parseInt(port), timeout }, () => {
            const status = `[UP] ${address} is reachable\nFrom: ${socket.localAddress}\nTo: ${socket.remoteAddress}`;
            resolve(green(status));
            socket.end();
        });

        socket.on("error", (err) => {
            const status = `[DOWN] ${address} is unreachable, error: ${err.message}`;
            resolve(red(status));
        });

        socket.on("timeout", () => {
            const status = `[DOWN] ${address} is unreachable, timed out.`;
            resolve(red(status));
            socket.end();
        });
    });
}

program
    .name("Healthy")
    .description("A CLI tool for checking the health of your services")
    .option("-d, --domain <domain>", "The domain to check", "localhost")
    .option("-p, --port <port>", "The port number to check", "80")
    .option("-l, --loop", "Loop the check every 5 seconds", false);

program.parse(process.argv);

const options = program.opts();
const { domain, port, loop } = options;

async function main() {
    if (loop) {
        while (true) {
            console.log(await performHealthCheck(domain, port));
            console.log("-".repeat(30));
            await new Promise((resolve) => setTimeout(resolve, 5000));
        }
    } else {
        console.log(await performHealthCheck(domain, port));
    }
}

main();
