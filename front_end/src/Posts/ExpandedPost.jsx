import {admin} from "../Static.js"
import { React, useState, useEffect } from "react";
import pPic from '../images/pp.png';
function ExpandedPost(props) {
    const postID = props.postID
    const exampleFriends = ["Kevin", "Omar" , "Raine", "Eugene"]
    const toggleHomepage = props.toggleHomepage
    const toggleOtherProfile = props.toggleOtherProfile
    const togglePost = () => props.toggleExpandPost(postID)

    const data = {
        userName: "Kevin",
        caption: "This is a post",
        image: "https://i.imgur.com/1w3o6Uo.jpg",
        comments: [
            {
                Commenter: "Kevins SQL key",
                Comment: "Way to Go!!",
                Profile: "https://i.imgur.com/1w3o6Uo.jpg",
            },
            {
                Commenter: "Victors SQL key",
                Comment: "Stop making Go puns",
                Profile: "https://i.imgur.com/1w3o6Uo.jpg",
            },
        ],
    };
    const [post, setPost] = useState(null);
    const [userComment, setuserComment] = useState("");
    const getPost = async () => {
        // ask Go
        setPost(data);
    };
    useEffect(() => {
        getPost();
    }, []);
    const makeComment = () => {};
    return (
    <div className="flex flex-col items-center">
        <button className="mb-5 w-fit ml-10 mt-5 mr-auto text-5xl hover:text-purple-300" onClick={toggleHomepage}> {'←'} </button>
        <div className="flex flex-col bg-white text-black text-start text-lg m-5 md:py-6 sm:px-16 lg:px-24 p-6 rounded-xl w-3/4 md:w-1/2 min-w-fit">
            <div className="flex flex-row justify-between items-center mt-8">
                <div className="flex flex-row justify-center items-center align-middle">
                    <img src="" alt="Profile Picture" className="w-10 aspect-square rounded-full"/>
                    <p className="ml-4">Eugene</p>
                </div>
                <p className="">Nov. 14th</p>
            </div>
            <p className="mt-10">This is the caption......</p>
            <div className="w-full bg-purple-200 rounded-lg p-2 my-2">
                <img src={pPic} className="my-4 mx-auto h-80 object-contain" alt="the post picture"/>
            </div>
            <div className="flex flex-col w-full bg-purple-200 rounded-xl p-2 my-5">
                <div className="w-full bg-purple-400 rounded-lg p-2 my-2">Add comment...</div>
                <div className="w-full bg-purple-400 rounded-lg p-2 my-2">I don't give up that easily</div>
                <div className="w-full bg-purple-400 rounded-lg p-2 my-2">Go to the Sun, it's got a warmer climate</div>
            </div>
        </div>
    </div>

        // <>
        //     {post == null ? null : (
        //         <div className="lg flex justify-center items-center relative border border-white border-solid bg-red-500">
        //             <button className="absolute top-0 left-0">&larr; Prev Page</button>
        //             <div className="md container left-12 relative">
        //                 <div>{post.userName}</div>

        //                 <div>{post.caption}</div>
        //                 <img src={post.image}></img>
        //                 <div>
        //                     <img src={post.image}></img>
        //                     <label>Add Comment</label>
        //                     <input
        //                         type="text"
        //                         value={userComment}
        //                         onChange={(e) => setuserComment(e.target.value)}
        //                         className="text-black"></input>
        //                 </div>
        //                 {userComment}
        //                 {post.comments.map((comment, index) => (
        //                     <div className="grid m-10" key={index}>
        //                         <img src={comment.Profile} className="aspect-square rounded-full"></img>
        //                         <div>{comment.Commenter}</div>
        //                         <div>{comment.Comment}</div>
        //                     </div>
        //                 ))}
        //             </div>
        //         </div>
        //     )}
        // </>
    );
}

export default ExpandedPost;
