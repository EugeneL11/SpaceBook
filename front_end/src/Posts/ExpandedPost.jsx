import currentUser from "../Static.js";
import { React, useState, useEffect } from "react";
import axios from 'axios'
import { serverpath } from "../Path.js";
function ExpandedPost(props) {
    const postID = props.post_id
    const toggleHomepage = props.toggleHomepage
    const toggleOtherProfile = props.toggleOtherProfile
    const togglePost = () => props.toggleExpandPost(postID)
    const [post, setPost] = useState(null);

    const [userComment, setUserComment] = useState(null);
    const [userCommentValue, setUserCommentValue] = useState("");
    const [numLikes, setNumLikes] = useState(0);

    useEffect(() => {
        // ask back end for post
        const path = `/postdetails/${encodeURIComponent(postID)}/${encodeURIComponent(currentUser.userID)}`
        axios.get(`${serverpath}${path}`).then((res) => {
            const data = res.data
            setPost(data.post)
            setUserComment(data.post.comments);
            setNumLikes(data.post.num_likes)
            if (data.post.liked) {
                setIsLiked(true)
            }
        })
    }, []);

    const handleKeyPress = (event) => {
        // Check if the Enter key was pressed (key code 13)
        if (event.key === 'Enter') {
            // Trigger the button click action
            makeComment();
        }
    };
    const makeComment = () => {
        const commentPath = `/makecomment/${encodeURIComponent(postID)}/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(userCommentValue)}`
        axios.post(`${serverpath}${commentPath}`).then((res) => {
            const data = res.data
            if (data.status === "no error") {
                const userCommentArr = userComment || [];
                const newArr = [{commenter_name: currentUser.userName, content: userCommentValue}, ...userCommentArr]
                setUserComment(newArr)
                setUserCommentValue("")
            } else {
                console.log(data.status)
            }
        })
    };
    
    //for admin
    const removePost = () =>{
        // ask back end
        const path = `/deletepost/${postID}`
        axios.delete(`${serverpath}${path}`).then((res) => {
            const data = res.data
            if (data.status !== "no error") {
                console.log(data)
            }
        })
        toggleHomepage();
    }

    const [isLiked, setIsLiked] = useState(false);
    
    const handleLike = () => {
        if (!isLiked) {
        setIsLiked(true);
        // Perform any additional actions when liked
            const path  = `/likepost/${encodeURIComponent(postID)}/${encodeURIComponent(currentUser.userID)}`
            axios.put(`${serverpath}${path}`).then(res =>{
                if (res.data.status === "no error") {
                    setNumLikes(numLikes + 1)
                    setIsLiked(true)
                } else {
                    console.log(data.status)
                }
            })
        }
    };

    const [imageNum,setImageNum] = useState(0)

    const toggleNextImage = () =>{
        let num = imageNum;
        if(num < post.images.length-1){
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

    return (
    <>
    {post === null ? <div></div> : (
        <div className="flex flex-col items-center">
            <div className="w-full flex items-center">
                <button className="mb-5 w-fit ml-6 mt-5 mr-auto text-3xl hover:text-purple-300" onClick={toggleHomepage}> {'←'} </button>
                {currentUser.admin && (<button className="mr-6 p-2 h-12 bg-red-200 hover:bg-red-400 rounded-md" onClick={removePost}>Delete Post</button>)} 
            </div>
            <div className="flex flex-col bg-white text-black text-start text-lg mb-10 md:py-6 sm:px-16 lg:px-24 p-6 rounded-xl w-3/4 lg:w-1/2 min-w-fit">
                <div className="flex flex-row pt-3 justify-between">
                    <div onClick={() => toggleOtherProfile(post.author_id, toggleHomepage)} className="flex flex-row justify-center align-middle pt-2 hover:cursor-pointer hover:text-purple-100">
                        <img src={serverpath + post.author_profile_path} alt="Profile Picture" className="w-10 aspect-square rounded-full"/>
                        <p className="ml-2 mt-2">{post.author_name}</p>
                    </div>
                    <p className="mr-2 mt-2 pr-2 pt-2">{post.date.substring(0, post.date.length - 10)}</p>
                </div>

                <p className="mt-10">{post.caption}</p>

                {
                post.images ?
                    post.images[imageNum] ?
                        <img src={serverpath + post.images[imageNum]} className="my-4 mx-auto h-48 object-contain" alt="the post picture"/>
                    : null
                : null
                }

            { post.images ?
            <div className="flex justify-center gap-10">
                {imageNum > 0 ? <button onClick={togglePrevImage} className="hover:text-gray-300"> Back </button> : null}
                {imageNum < post.images.length - 1 ? <button onClick={toggleNextImage} className="hover:text-gray-300"> Next </button> : null}
            </div>
            : null }

                <div className="flex flex-col w-full bg-purple-200 rounded-xl p-2 my-5">
                    <div className="flex justify-between">
                        <div>{numLikes} likes</div>
                        <div
                            className="cursor-pointer transition duration-300 ease-in-out"
                            onClick={handleLike}
                            >
                            <img
                                src={isLiked ? 'redHeart.png' : 'blackHeart.png'}
                                alt="Heart"
                                className="h-8 w-8"
                            />
                        </div>
                    </div>
                    <div className="flex w-full bg-purple-300 rounded-lg p-2 my-2">
                        <input  
                            className="w-full bg-transparent border-b-2 border-gray-600 focus:outline-none focus:border-gray-300 focus:ring-0 text-black placeholder-gray-500"
                            placeholder="Add comment..." 
                            type="text"
                            value = {userCommentValue}
                            onKeyPress={userCommentValue ? handleKeyPress : null}
                            onChange = {(e) => {setUserCommentValue(e.target.value)}}
                        >
                        </input>
                        <button className="p-2 bg-blue-300 hover:bg-blue-400 text-white rounded-md ml-2" onClick={userCommentValue ? makeComment : null}><img src="arrow-up.png" className="w-4"></img></button>
                    </div>
                    {userComment ? userComment.map((comment, index) => (
                        <div key = {index}>
                            <div className="w-full flex justify-between bg-purple-400 rounded-lg p-2 my-2">
                                <div onClick ={() => toggleOtherProfile(userComment.commenter_id, togglePost)} className="text-white mr-6 text-left">{comment.commenter_name}</div>
                                <div className="h-14 overflow-y-scroll text-right">{comment.content}</div>
                            </div>
                        </div>)
                    ) : null}
                </div>
                {(currentUser.userID === post.author_id) ? 
                    <button className="mx-auto p-2 w-fit h-12 bg-red-200 hover:bg-red-400 rounded-md" onClick={removePost}>Delete Post</button> : 
                    null
                }
            </div>
        </div>
    )}
    </>
    )
}

export default ExpandedPost;