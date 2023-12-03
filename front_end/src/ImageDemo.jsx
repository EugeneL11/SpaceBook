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
        console.log(newImages)
    };

    const saveEvent = imageUpload

    const toggleNextImage = () =>{
        let num = imageNum;
        if(num < images.length-1){
            num++;
        }
        console.log(num)

        setImageNum(num);
        handleImageChange({target: {files: [images[num]]}})
    }

    const togglePrevImage = () =>{
        let num = imageNum;
        if(num > 0){
            num--;
        }
        console.log(num)

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
        console.log(selectedImage)

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


            {/* { imageNum > 0 ? 
                <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer translate-x-5 -translate-y-16 rotate-180" onClick={togglePrevImage} /> : null
            } */}

            <img src= {previewImage ? previewImage: defaultIMG} className="my-4 mx-auto h-40  object-contain"></img>


            {/* {imageNum < images.length - 1 ? 
            <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer right-0 -translate-x-5 -translate-y-16"  onClick={toggleNextImage}/> : null
            } */}

            </div>
            <div className="flex justify-center gap-10">
                <button onClick={togglePrevImage} className="hover:text-gray-300"> Back </button>
                <button onClick={toggleNextImage} className="hover:text-gray-300"> Next </button>
            </div>

            {/* <button onClick={uploadTest}>Send post message to server</button> */}
        </div>
    );
}

export default ImageDemo;