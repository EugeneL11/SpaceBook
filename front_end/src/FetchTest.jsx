import { React, useState } from "react";

function FetchTest() {
    const [sum, setSum] = useState(0);
    fetch("http://localhost:8080/postgresTest").then((res) => {
        res.json().then((result) => {
            setSum(result.value);
        });
    });
    return <div>The sum of two numbers computed by our backend: {sum}</div>;
}

export default FetchTest;
