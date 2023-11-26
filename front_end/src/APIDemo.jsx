import { React, useState , useEffect } from "react";
import axios from 'axios'
function APIDemo(){

    // the two below have the same behavior
    const aysncReq = async() =>{
        try{
            // This is the raw binary data the server sent you
            const json_file = await fetch("http://localhost:8080/ping")

            // We need to use .json() which is also async to convert it to an object or array
            const javascript_object = json_file.json()

            // Do stuff with javascript object, probably set state
        }
        catch{
            console.log("Error or something idk")
        }
    }
    const sequentialRequest = () => {
        // .then(json_file => {}) means after this is done, execute this lamba function
        // json_file is the result returned from the async
        fetch("http://localhost:8080/ping").then(json_file => {
            // same as above .then() takes in a lamba, javascript_object is the returned value
            json_file.json().then(javascript_object =>{
                // Do stuff with javascript object, probably set state
            })
        }).catch(
            // Maybe the server didn't respond
        )
    }

    return (<div>Go Go Go</div>)
}