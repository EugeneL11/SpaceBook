import {React, useState} from "react";

function FetchTest(){
    
    const [sum, setSum] = useState(0);
    //fetch("http://localhost:3001/2/1").then(res => {res.json().then(result => {setSum(result)}})
    return(
        <div>The sum of two numbers computed by our backend: {sum}</div>
    );
}

export default FetchTest;
