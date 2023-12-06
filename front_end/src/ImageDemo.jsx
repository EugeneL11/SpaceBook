import React, { useState, useRef } from "react";
import axios from 'axios';
import { serverpath } from "./Path";

function ImageDemo(props) {
    const updateImage = props.setImages;

    const [image, setImage] = useState(null);
    const [previewImage, setPreview] = useState(null)
    const defaultIMG = (`${serverpath}/images/header.jpg`)

    const [imageNum,setImageNum] = useState(0)

    const [images, setImages] = useState([])

    const imageUpload = (file) => {
        const newImages = [...images, file]
        setImages(newImages)
        updateImage(newImages)
    };

    const saveEvent = imageUpload

    const toggleNextImage = () =>{
        let num = imageNum;
        if(num < images.length-1){
            num++;
        }

        setImageNum(num);
        handleImageChange({target: {files: [images[num]]}})
    }

    const togglePrevImage = () =>{
        let num = imageNum;
        if(num > 0){
            num--;
        }

        setImageNum(num);
        handleImageChange({target: {files: [images[num]]}})
    }

    const imageIsUploaded = (e) => {
        const selectedImage = e.target.files[0];
        setImage(selectedImage);
        saveEvent(selectedImage)

        handleImageChange(e);
    }

    const handleImageChange = (e) => {
        const selectedImage = e.target.files[0];

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
            <input type="file" onChange={imageIsUploaded} />

            <div className="relative w-100 h-100 mt-4 border-black border-2">

            <img src= {previewImage ? previewImage: defaultIMG} className="my-4 mx-auto h-40  object-contain"></img>

            </div>
            <div className="flex justify-center gap-10">
                {imageNum > 0 ? <button onClick={togglePrevImage} className="hover:text-gray-300"> Back </button> : null}
                {imageNum < images.length - 1 ? <button onClick={toggleNextImage} className="hover:text-gray-300"> Next </button> : null}
            </div>

            {/* <button onClick={uploadTest}>Send post message to server</button> */}
        </div>
    );
}

export default ImageDemo;