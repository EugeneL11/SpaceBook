import React, { useState } from "react";
import axios from 'axios';
import { serverpath } from "./Path";

function ImageDemo(props) {

    const [image, setImage] = useState(null);
    const [previewImage, setPreview] = useState(null)
    const saveEvent = props.saveEvent
    const defaultIMG = (props.image ? props.image : `${serverpath}/images/header.jpg`)
    // change these
    const localhost = "http://localhost:8080";
    const path = `/uploadpostimage/${encodeURIComponent("b86e1293-8bee-11ee-b365-0242ac120004")}`;

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
    const handleImageChange = (e) => {
        const selectedImage = e.target.files[0];
        setImage(selectedImage);
        saveEvent(selectedImage)
        // Read and set the preview image
        const reader = new FileReader();
        reader.onloadend = () => {
            setPreview(reader.result);
        };
        if (selectedImage) {
            reader.readAsDataURL(selectedImage);
        }
    };
    return (
        <div>
            <h1>Click Below to upload image</h1>
            <input type="file" onChange={handleImageChange} />
            <img src= {previewImage ? previewImage: defaultIMG}></img>
            
            <img src={props.image} alt="" />
            <button onClick={uploadTest}>Send post message to server</button>
        </div>
    );
}

export default ImageDemo;
