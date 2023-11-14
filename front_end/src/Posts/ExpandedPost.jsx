import {admin} from "../Static.js"
import { React, useState, useEffect } from "react";
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
