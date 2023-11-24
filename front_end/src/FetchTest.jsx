import { React, useState } from "react";

function FetchTest() {
    const [sum, setSum] = useState(0);


    const user = {
        ID: "1",
        Name: "John",
    };

    // fetch("http://localhost:8080/postgresTest").then((res) => {
    //     res.json().then((result) => {
    //         setSum(result.value);
    //     });
    // });
    // fetch("http://localhost:8080/user").then((res) => {
    //     res.json().then((result) => {
    //         setSum(result.value);
    //     });
    // });

    // fetch("http://localhost:8080/ping", {
    //     method: "POST",
    //     body: JSON.stringify(user),
    // }).then((res) => {
    //     res.json().then((result) => {
    //         setSum(result.value);
    //     });
    // });

    fetch("https://space-book-pied.vercel.app/ping").then((res) => {
        res.json().then((result) => {
            setSum(result.message);
        });
    });

    return <div>The sum of two numbers computed by our backend: {sum}</div>;
}

export default FetchTest;
