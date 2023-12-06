import { React, useState, useRef } from "react";
import axios from 'axios'
import currentUser from "../Static";
import ImageDemo from "../ImageDemo";


// import {imageNum, togglePrevImage, toggleNextImage, images} from "../ImageDemo"


import { serverpath } from "../Path";

function NewPost(props) {
    const toggleHomepage = props.toggleHomepage

    const examplePost = {
        username: "Duppy",
        pfp : "./ayylmao.webp",
        caption: "Finally leaving this planet lmao 😂",
        date: "Nov 7th",
        images: ["./swag.jpg", "./ayylmao.webp"],
        videos: [],
    }


   // function makePost() {

    const [imageNum,setImageNum] = useState(0)
    const [images, setImages] = useState([null,null,null,null,null])
    const [selectedImage, setSelectedImage] = useState(null)
    const fileInputRef = useRef(null)
    const [caption, setCaption] = useState(' ')
    const [postID, setPostID] = useState(0)
   
    const handleCaption = (event) => {
      setCaption(event.target.value);
    };


    async function makePost() {
        const path = `/makepost/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(caption)}`
        const res = await axios.post(`${serverpath}${path}`)
        const data = res.data
        if (data.status === "no error") {
            setPostID(data.post_id)
        } else {
            console.log(data.status)
        }
        for (let i = 0; i < images.length; i++) {
            if (images[i] !== null) {
                const formData = new FormData();
                formData.append(`image`, images[i])
                const imagePath = `/uploadpostimage/${encodeURIComponent(data.post_id)}`
                const imageRes = await axios.post(`${serverpath}${imagePath}`, formData)
                const imageData = imageRes.data
            }
        }
        toggleHomepage();
        //take to homepage
    }
    
    return (
        <div className="flex flex-col items-center">
            <button className="mb-5 w-fit ml-6 mr-auto text-3xl hover:text-purple-300" onClick={toggleHomepage}> {'←'} </button>

            <div className="flex flex-col bg-white text-black text-start text-lg mt-[-20px] mb-10 md:py-6 sm:px-16 lg:px-24 p-6 rounded-xl w-3/4 md:w-1/2 min-w-fit">
                <div className="flex justify-center items-center">
                    <h1 className="text-center text-3xl pt-2">New Post</h1>
                </div>

                <label className=" mt-3 lg:mt-4">Write a caption </label>
                <textarea 
                    className="form-textarea border-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" 
                    rows="3" placeholder=" Write a caption..."
                    value={caption} onChange={handleCaption}    
                ></textarea>
            
                <div className="mt-4">Add images </div>
          

                <h1>Click Below to upload image</h1>

                    <ImageDemo
                        setImages = {setImages}
                    /> 

                <button onClick={makePost} className="bg-purple-300 hover:bg-purple-400 px-5 py-2 mt-5 w-fit self-center rounded-lg">Post!</button>

            </div>

        </div>
    ); 
}

export default NewPost;


// {imageNum > 0 ? 
//     // <div className="absolute text-purple-500 pb-2 pr-2 pl-2 bg-slate-300 bg-opacity-60 rounded-full text-7xl top-52 z-40 cursor-pointer hover:text-purple-400" onClick={togglePrevImage}> {"←"} </div> : null
    // <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer translate-x-5 -translate-y-16 rotate-180" onClick={togglePrevImage} /> : null
// }

// {/* <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer translate-x-5 -translate-y-16 rotate-180" onClick={toggleHomepage} />  */}
// <img src={examplePost.images[imageNum]} className="h-80 object-contain ml-auto mr-auto" alt="" />
// {/* <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer right-0 -translate-x-5 -translate-y-16"  onClick={toggleHomepage}/> */}


// {imageNum < imageCount - 1 ? 
// // <div className="absolute text-purple-500 pb-2 pr-2 pl-2 bg-slate-300 bg-opacity-60 rounded-full text-7xl cursor-pointer left-85-percent top-52 z-40 hover:text-purple-400" onClick={toggleNextImage}> {"→"} </div> : null
// <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer right-0 -translate-x-5 -translate-y-16"  onClick={toggleNextImage}/> : null
// }