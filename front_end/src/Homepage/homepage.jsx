import { React, useState, useEffect } from "react";
import {userID} from "../Static.js"

function Post(props) {
    const imageCount = props.post.images.length;
    const toggleOtherProfile = props.toggleOtherProfile;
    const expandPost = props.expandPost;
    const [imageNum,setImageNum] = useState(0)
    const toggleNextImage = () =>{
        const nextImage = imageNum + 1;
        setImageNum(nextImage);
    }
    const togglePrevImage = () =>{
        const nextImage = imageNum - 1;
        setImageNum(nextImage);
    }
    return(
        <div onClick={expandPost} className="flex flex-col w-5/12 mx-auto mb-10 justify-center align-middle bg-slate-300 bg-opacity-90 text-black rounded-lg">
            <div className="relative w-100 h-100">
            {imageNum > 0 ? 
                <div className="absolute text-purple-500 pb-2 pr-2 pl-2 bg-slate-300 bg-opacity-60 rounded-full text-7xl top-52 z-40 cursor-pointer hover:text-purple-400" onClick={togglePrevImage}> {"←"} </div> : null}
            {imageNum < imageCount - 1 ? 
                <div className="absolute text-purple-500 pb-2 pr-2 pl-2 bg-slate-300 bg-opacity-60 rounded-full text-7xl cursor-pointer left-85-percent top-52 z-40 hover:text-purple-400" onClick={toggleNextImage}> {"→"} </div> : null}
                </div>
            <div className="flex flex-row p-3 justify-between">
                <div onClick={toggleOtherProfile} className="flex flex-row justify-center align-middle pl-2 pt-2 ">
                    <img src={props.post.pfp} alt="Profile Picture" className="w-10 aspect-square rounded-full"/>
                    <p className="ml-2 mt-2">{props.post.username}</p>
                </div>
                <p className="mr-2 mt-2 pr-2 pt-2">{props.post.date}</p>
            </div>

            <p className="mt-2 pl-5"> {props.post.caption}</p>
            {/* no resize on image */}
            <img src={props.post.images[imageNum]} className="m-4 h-80  object-contain " alt="" />
        </div>
    )
}

function Homepage(props) {
    
    const toggleExpandPost = props.toggleExpandPost;
    const toggleOtherProfile = props.toggleOtherProfile;
    const toggleHomePage = props.toggleHomePage;

    const [posts, setPosts] = useState(null);
    const examplePost = {
        username: "Duppy",
        pfp : "./ayylmao.webp",
        caption: "Finally leaving this planet lmao 😂",
        date: "Nov 7th",
        images: ["./swag.jpg", "./ayylmao.webp"],
        videos: [],
    };
    const examplePosts = [examplePost, examplePost];

    useEffect(() => {
         // ask back end for post
         setPosts(examplePosts);
    }, [])


    return (<div className="flex flex-col mt-5 pt-5"> 
        {
            posts ? posts.map((post,index) => 
                (<Post key = {index} post = {post}/>)
            ) : null
        }
    </div>);

}

export default Homepage;
