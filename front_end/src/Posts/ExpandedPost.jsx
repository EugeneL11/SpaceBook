import static1 from "../Static.js";
import { React, useState, useEffect } from "react";
import pPic from '../images/pp.png';
function ExpandedPost(props) {
    const postID = props.postID
    const exampleFriends = ["Kevin", "Omar" , "Raine", "Eugene"]
    const toggleHomepage = props.toggleHomepage
    const toggleOtherProfile = props.toggleOtherProfile
    const togglePost = () => props.toggleExpandPost(postID)

    const [post, setPost] = useState(null);
    const examplePost = {
        postId: 5,
        authorId: 6,
        username: "Duppy",
        pfp : "./ayylmao.webp",
        caption: "Finally leaving this planet lmao 😂",
        date: "Nov. 7th",
        images: ["./swag.jpg", "./ayylmao.webp"],
        comments: [{username: "duppy", content: "I don't give up that easily", id: 5}, {username: "kevon", content: "Go to the Sun, it's got a warmer climate", id: 6}]
    };

    useEffect(() => {
         // ask back end for post
         setPost(examplePost);
    }, []);

    // const imageCount = post.images.length;
    // const [imageNum,setImageNum] = useState(0)
    // const toggleNextImage = () =>{
    //     let nextImage = imageNum + 1;
    //     setImageNum(nextImage);
    // }
    // const togglePrevImage = () =>{
    //     let nextImage = imageNum - 1;
    //     setImageNum(nextImage);
    // }

    const [userComment, setuserComment] = useState("");
    const makeComment = () => {
        // ask backend
    };
    
    //for admin
    const removePost = () =>{
        // ask back end
        toggleHomepage();
    }
    
    return (
    <>
    {post === null ? <div></div> : (
        <div className="flex flex-col items-center">
            <div className="w-full flex items-center">
                <button className="mb-5 w-fit ml-10 mt-5 mr-auto text-5xl hover:text-purple-300" onClick={toggleHomepage}> {'←'} </button>
                {static1.admin && (<button className="mr-10 p-2 h-12 bg-red-200 hover:bg-red-400 rounded-md" onClick={removePost}>Delete Post</button>)} 
            </div>
            <div className="flex flex-col bg-white text-black text-start text-lg m-5 md:py-6 sm:px-16 lg:px-24 p-6 rounded-xl w-3/4 md:w-1/2 min-w-fit">
                {/* <div className="relative w-100 h-100">
                    {imageNum > 0 ?
                        <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer translate-x-10 translate-y-10 rotate-180" onClick={togglePrevImage} /> : null
                    }
                    {imageNum < imageCount - 1 ? 
                        <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer right-0 -translate-x-10 translate-y-10"  onClick={toggleNextImage}/> : null
                    }        
                </div> */}
                <div className="flex flex-row justify-between items-center mt-8">
                    <div className="flex flex-row justify-center items-center align-middle">
                        <img src={post.pfp} alt="Profile Picture" className="w-10 aspect-square rounded-full"/>
                        <p className="ml-4">{post.username}</p>
                    </div>
                    <p className="">{post.date}</p>
                </div>
                <p className="mt-10">{post.caption}</p>
                <div className="w-full bg-purple-200 rounded-lg p-2 my-2">
                    <img src={post.images[0]} className="my-4 mx-auto h-80 object-contain" alt="the post picture"/>
                </div>
                <div className="flex flex-col w-full bg-purple-200 rounded-xl p-2 my-5">
                    <div className="flex w-full bg-purple-300 rounded-lg p-2 my-2">
                        <input  
                            class="w-full bg-transparent border-b-2 border-gray-600 focus:outline-none focus:border-gray-300 focus:ring-0 text-black placeholder-gray-500"
                            placeholder="Add comment..."
                            type="text"
    //                        value = {messageValue}
    //                        onChange = {(e) => {setmessageValue(e.target.value)}}
                            >
                        </input>
                        <button className="p-2 bg-blue-300 hover:bg-blue-400 text-white rounded-md ml-2" onClick={makeComment}><img src="arrow-up.png" className="w-4"></img></button>
                    </div>

                    {post.comments.map((comment ,index) => (
                        <div key = {index}>
                            <div className="w-full flex justify-between bg-purple-400 rounded-lg p-2 my-2">
                                <div className="text-white">{comment.username}</div>
                                <div>{comment.content}</div>
                            </div>
                        </div>)
                    )}
                </div>
            </div>
        </div>
    )}
    </>
    )
}

export default ExpandedPost;
