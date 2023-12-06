import { React, useState, useEffect } from "react";
import currentUser from "../Static.js"
import axios from 'axios'
import { serverpath } from "../Path.js";

function Post(props) {
//    const imageCount = props.images.length;
    const toggleOtherProfile = props.toggleOtherProfile;
    const toggleExpandPost = props.toggleExpandPost;
    const toggleHomePage = props.toggleHomePage;
    const expandPost = props.expandPost;
    
    const [imageNum,setImageNum] = useState(0)

    const toggleNextImage = () =>{
        let num = imageNum;
        if(num < props.images.length-1){
            num++;
        }
        setImageNum(num);
    }

    const togglePrevImage = () =>{
        let num = imageNum;
        if(num > 0){
            num--;
        }
        setImageNum(num);
    }
    return(
        <div className="flex flex-col bg-white text-black text-start text-lg mx-auto mb-10 md:py-6 sm:px-16 lg:px-24 p-6 rounded-xl w-3/4 min-w-fit">
            
            <div className="flex flex-row pt-3 justify-between">
                <div onClick={() => toggleOtherProfile(props.authorID, toggleHomePage)} className="flex flex-row justify-center align-middle pl-2 pt-2 hover:cursor-pointer hover:text-purple-100">
                    <img src={props.authorPP} alt="Profile Picture" className="w-10 aspect-square rounded-full"/>
                    <p className="ml-2 mt-2">{props.authorName}</p>
                </div>
                <p className="mr-2 mt-2 pr-2 pt-2">{props.date.substring(0, props.date.length - 10)}</p>
            </div>

            <p className="mt-10"> {props.caption}</p>
            {/* no resize on image */}
            {/* <img src={props.images[imageNum]} className="m-4 h-80  object-contain " alt="" /> */}
            {/* {(props.images && props.images.length > 0) ? (
                <div className="flex flex-wrap gap-4 mx-auto">
                {props.images.map((image, index) => (
                    <img
                    key={index}
                    src={serverpath+image}
                    alt={`Image ${index + 1}`}
                    className="max-w-full h-40 rounded-lg"
                    />
                ))}
                </div>
            ) : null} */}

            {
                props.images ?
                
                    props.images[imageNum] ?
                        <img src={serverpath + props.images[imageNum]} className="my-4 mx-auto h-48 object-contain" alt="the post picture"/>
                    : null

                : null
            }

            { props.images ?
            <div className="flex justify-center gap-10">
                {imageNum > 0 ? <button onClick={togglePrevImage} className="hover:text-gray-300"> Back </button> : null}
                {imageNum < props.images.length - 1 ? <button onClick={toggleNextImage} className="hover:text-gray-300"> Next </button> : null}
            </div>
            : null }

            <button onClick={toggleExpandPost} className="bg-purple-300 hover:bg-purple-400 px-7 py-3 m-6 w-fit self-center rounded-lg">Expand Post</button>
        </div>
    )
}

function Homepage(props) {
    
    const toggleExpandPost = props.toggleExpandPost;
    const toggleOtherProfile = props.toggleOtherProfile;
    const toggleHomePage = props.toggleHomePage;
    const [posts, setPosts] = useState(null);

    useEffect(() => {
        // ask back end for post
        const path = `/homepageposts/${encodeURIComponent(currentUser.userID)}`
        axios.get(`${serverpath}${path}`).then((res) => {
            const data = res.data
            if (data.status === "no error") {
                setPosts(data.posts)
            } else {
                console.log(data.status)
            }
        })
    }, [])

    return (
    <div className="flex flex-col mt-5 pt-5"> 
        {
            posts ? posts.map((post,index) => 
                (<Post 
                    key = {index} post = {post} 
                    authorID = {post.author_id} authorName = {post.author_name}
                    authorPP = {serverpath + post.author_profile_path} images = {post.images}
                    date = {post.date} caption = {post.caption}
                    toggleExpandPost = {() =>{toggleExpandPost(post.post_id)}} 
                    toggleHomePage = {toggleHomePage} 
                    toggleOtherProfile = {toggleOtherProfile}/>)
            ) : <NoPosts/>
        }
    </div>
    );
}

function NoPosts() {
    return(
        <div className="flex flex-col mt-5 w-fit max-w-[85%] bg-white rounded-lg text-black text-center text-xl mx-auto p-10"> 
            <p className="text-3xl text-center">No Posts Yet.</p>
            <p className="text-3xl text-center">To see posts, follow another user and have them make a post!</p>
        </div>
    )
}

export default Homepage;