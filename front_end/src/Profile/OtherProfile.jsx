import { React, useState, useEffect } from "react";
import Planet from "./Planet.jsx";
import { Canvas, useThree } from "@react-three/fiber";
import axios from 'axios'
import currentUser from "../Static";
import {serverpath} from "../Path.js";
import * as THREE from 'three';

function ResizingCanvas(props) {
    const { gl, size, camera } = useThree();
    var glsize

    //resizing the planets to make it responsive
    useEffect(() => {
        const handleResize = () => {
            // Update size
            if (window.innerWidth < 1024) {
                glsize = 150
            } else {
                glsize = 270
            }

            gl.setSize(glsize, glsize);
            camera.aspect = 1;
            camera.updateProjectionMatrix();
        };
        window.addEventListener('resize', handleResize);

        // Clean up on unmount
        return () => window.removeEventListener('resize', handleResize);
    }, [gl, camera]);
    return null;
}

// each planet component
function PlanetCanvas (props) {
    return (
        <Canvas className="cursor-pointer pt-5 lg:pt-0">
            <Planet planet={props.planet}/>
            <ResizingCanvas className=" translate-x-full"/>
        </Canvas>
    );
}

function Post(props) {
    const expandPost = props.expandPost;
    const [imageNum,setImageNum] = useState(0)
    
    // move on to the next image
    const toggleNextImage = () =>{
        const nextImage = imageNum + 1;
        setImageNum(nextImage);
    }

    // move back to the previous image
    const togglePrevImage = () =>{
        const nextImage = imageNum - 1;
        setImageNum(nextImage);
    }
    
    // html code for a single post
    return(
        <div className="flex flex-col w-11/12 lg:w-7/12 xl:w-5/12 mx-auto mb-10 md:py-6 sm:px-16 lg:px-20 p-6 justify-center align-middle bg-slate-300 bg-opacity-90 text-black rounded-lg">
            <div className="flex flex-row pt-3 justify-between">
                <div className="flex flex-row justify-center align-middle pt-2">
                    <img src={serverpath + props.post.author_profile_path} alt="Profile Picture" className="w-10 aspect-square rounded-full"/>
                    <p className="ml-2 mt-2">{props.post.author_name}</p>
                </div>
                <p className="mr-2 mt-2 pr-2 pt-2">{props.post.date.substring(0, props.post.date.length - 10)}</p>
            </div>

            <p className="mt-10"> {props.post.caption}</p>

            {props.post.images 
                ? props.post.images[imageNum] 
                    ? <img src={serverpath + props.post.images[imageNum]} className="my-4 mx-auto h-48 object-contain" alt="the post picture"/>
                    : null
                : null
            }

            { props.post.images ?
                <div className="flex justify-center gap-10">
                    {imageNum > 0 ? <button onClick={togglePrevImage} className="hover:text-gray-300"> Back </button> : null}
                    {imageNum < props.post.images.length - 1 ? <button onClick={toggleNextImage} className="hover:text-gray-300"> Next </button> : null}
                </div>
            : null }

            <button onClick={expandPost} className="bg-purple-300 hover:bg-purple-400 px-7 py-3 m-6 w-fit self-center rounded-lg">Expand Post</button>
        </div>
    )
}

function OtherProfile(props) {
    const back = props.goBackScreen
    const personID = props.userID
    const togglePost = props.togglePost
    const [user, setUser] = useState(null)
    const [posts, setPosts] = useState(null)
    const [friendStatus, setFriendStatus] = useState("")
        
    // ask back end for user information
    useEffect(() => {
        const path = `/getuserinfo/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(personID)}`
        axios.get(`${serverpath}${path}`).then(res => {
            const data = res.data
            setUser(data.user)
            setFriendStatus(data.friendstatus)
            setPosts(data.posts)
        })
    }, [])

    // tell backend to delete the user
    const removeUser = () =>{
        const path = `/deleteuser/${encodeURIComponent(personID)}`
        axios.delete(`${serverpath}${path}`).then(res => {
            if (res.data.status == "no error") {
                back()
            }
            else {
                console.log(res.data.status)
            }
        })
    }

    // tell backend to send a friend request
    function sendOrbit() {
        const path = `/sendfriendreq/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(props.userID)}`
        axios.post(`${serverpath}${path}`).then(res => {
            const data = res.data
            if (data.status === "no error") {
                setFriendStatus("viewer sent request")
            } else {
                console.log(data.status)
            }
        })
    }

    // tell the backend to remove user as friend
    function unorbit() {
        const path = `/removefriend/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(props.userID)}`
        axios.delete(`${serverpath}${path}`).then(res => {
            const data = res.data
            if (data.status === "no error") {
                setFriendStatus("no requests")
            } else {
                console.log(data.status)
            }
        })
    }

    // tell backend to reject the user's friend request
    function rejectOrbitRequest() {
        const path = `/rejectfriendreq/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(props.userID)}`
        axios.delete(`${serverpath}${path}`).then(res => {
            const data = res.data
            if (data.status === "no error") {
                setFriendStatus("no requests")
            } else {
                console.log(data.status)
            }
        })
    }

    // the backend to accept the user's friend request
    function acceptFriendRequest() {
        const path = `/sendfriendreq/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(props.userID)}`
        axios.post(`${serverpath}${path}`).then(res => {
            const data = res.data
            if (data.status === "no error") {
                setFriendStatus("already friends")
            } else {
                console.log(data.status)
            }
        })
    }

    // all the different scenarios of the friendship status
    const table = {
        "" : null,
        "already friends": <div className = "mb-6 hover:text-red-300" onClick={unorbit}>Unorbit this buddy</div>,
        "no requests": <div className = "mb-6 hover:text-green-300" onClick={sendOrbit}>Request orbit buddy</div>,
        "viewer sent request": <div>Orbit request pending</div>,
        "viewed person sent request": <div className="flex flex-col">
            <div className = "mb-1 hover:text-green-300" onClick={acceptFriendRequest}>Accept orbit request</div>
            <div className = "mb-6 hover:text-red-300" onClick={rejectOrbitRequest}>Reject orbit request</div>
        </div>
    }

    // html code displaying a user's profile
    return (
        <div className="flex flex-col">
        <div className="w-full flex items-center">
            <button className="mb-5 w-fit ml-6 mr-auto text-3xl hover:text-purple-300" onClick={back}> {'←'} </button>
            {currentUser.admin && (<button className="mr-6 p-2 h-12 bg-red-200 hover:bg-red-400 rounded-md" onClick={removeUser}>Remove User</button>)} 
        </div>

        { user ? 
        
        <div className="flex flex-col justify-center align-middle">

        <div className="flex flex-col lg:flex-row justify-center align-middle mb-14">
    
            <div className="flex flex-col ml-5 mr-5 lg:ml-0 lg:mr-0 lg:w-80">

                <div className="flex flex-row mb-4 align-middle">

                    <img src={serverpath + user.profile_picture_path} alt="My Profile Picture" className="w-20 aspect-square rounded-full"/>

                    <div className="flex flex-col ml-4 justify-center align-middle">
                        <p className="text-xl">{user.user_name}</p>
                        <p>from {user.planet}</p>
                    </div>

                </div>

                <p className="mb-4">{user.bio}</p>

                <div className="flex">
                    <img src="./addwhite.png" className="h-5 aspect-square translate-y-0.5 mr-2"/>
                    { table[friendStatus] }
                </div>

            </div>

            <div className="translate-x-1/3 w-2/3 pl-7 md:pl-36 lg:pl-0 lg:translate-x-0 lg:w-auto">
                <PlanetCanvas planet={user.planet}/>
            </div>

            </div>
            {
                posts ? posts.map((post,index) => 
                (<Post key = {index} post = {post} expandPost ={() => togglePost(post.post_id)}/>)
                ) :
                <div className="w-fit bg-white rounded-lg text-black text-center text-xl mx-auto p-10">
                    This User Has Yet To Post...
                </div> 
            }

        </div>: null}

    </div>
        
    );
}

export default OtherProfile;