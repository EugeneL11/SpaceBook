import React, { useState } from "react";
import axios from 'axios';

function ImageDemo() {
    const [image, setImage] = useState(null);

    // change these
    const localhost = "https://space-book-pied.vercel.app";
    const path = "/upload";

    const uploadTest = () => {
        const formData = new FormData();
        formData.append("image", image);

        axios.post(`${localhost}${path}`, formData)
            .then(response => {

                console.log(response);
            })
            .catch(error => {

                console.error(error);
            });
    };

    return (
        <div>
            <h1>Click Below to upload image</h1>
            <input type="file" onChange={(e) => setImage(e.target.files[0])} />
            <button onClick={uploadTest}>Send post message to server</button>
        </div>
    );
}

export default ImageDemo;
