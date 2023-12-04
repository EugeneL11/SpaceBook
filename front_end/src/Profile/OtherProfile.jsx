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


function PlanetCanvas (props) {

    return (
        <Canvas className="cursor-pointer pt-5 lg:pt-0">
            <Planet planet={props.planet}/>
            <ResizingCanvas className=" translate-x-full"/>
        </Canvas>
    );
}

function Post(props) {
   //const imageCount = props.post.images.length;
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
        <div className="flex flex-col w-11/12 lg:w-7/12 xl:w-5/12 mx-auto mb-10 md:py-6 sm:px-16 lg:px-20 p-6 justify-center align-middle bg-slate-300 bg-opacity-90 text-black rounded-lg">
            {/* <div className="relative w-100 h-100">
                {imageNum > 0 ? 
                    // <div className="absolute text-purple-500 pb-2 pr-2 pl-2 bg-slate-300 bg-opacity-60 rounded-full text-7xl top-52 z-40 cursor-pointer hover:text-purple-400" onClick={togglePrevImage}> {"←"} </div> : null
                    <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer translate-x-10 translate-y-10 rotate-180" onClick={togglePrevImage} /> : null
                }
                {imageNum < imageCount - 1 ? 
                    // <div className="absolute text-purple-500 pb-2 pr-2 pl-2 bg-slate-300 bg-opacity-60 rounded-full text-7xl cursor-pointer left-85-percent top-52 z-40 hover:text-purple-400" onClick={toggleNextImage}> {"→"} </div> : null
                    <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer right-0 -translate-x-10 translate-y-10"  onClick={toggleNextImage}/> : null

                }
            </div> */}
            <div className="flex flex-row pt-3 justify-between">
                <div className="flex flex-row justify-center align-middle pt-2">
                    <img src={serverpath + props.post.author_profile_path} alt="Profile Picture" className="w-10 aspect-square rounded-full"/>
                    <p className="ml-2 mt-2">{props.post.author_name}</p>
                </div>
                <p className="mr-2 mt-2 pr-2 pt-2">{props.post.date.substring(0, props.post.date.length - 10)}</p>
            </div>

            <p className="mt-10"> {props.post.caption}</p>
            {/* no resize on image */}
            {/* {props.images ? 
                <img src={props.images[imageNum]} className="m-4 h-80  object-contain " alt="" />
            : null} */}

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
        </div>
    )
}

function OtherProfile(props) {
    const toggleFriendsList = props.toggleFriendsList
    const toggleSettings = props.toggleSettings
    const dm = props.goDMList
    // const dm = DMController(toggleHomepage = props.toggleHomepage )
    // const dm = setScreen(<DMController toggleHomepage={props.toggleHomepage} />);
    const back = props.goBackScreen

    const duppy = {
        username: "Rainethhh", 
        homePlanet: "Pluto👶",
        pfp: "./swag.jpg",
        bio: "What's up guys I'm Raine, I like hiking and coding and being chill+smart also I made a cool game ask me about it"
    }
    const [user, setUser] = useState(null)
    const [posts, setPosts] = useState(null)
    const examplePost = {
        username: "Rainethhh",
        pfp : "./swag.jpg",
        caption: "Just hiked up the second tallest mountain on Jupiter!",
        date: "Nov 9th",
        images: ["./swag.jpg", "./ayylmao.webp"],
        videos: [],
    }
    const examplePosts =[ examplePost,examplePost]

    const [friendStatus, setFriendStatus] = useState("")

    useEffect(() => {
        // ask bak end for user
        const path = `/getuserinfo/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(props.userID)}`
        axios.get(`${serverpath}${path}`).then(res => {
            const data = res.data
            console.log(currentUser.userID)
            console.log(data)
            setUser(data.user)
            setFriendStatus(data.friendstatus)
            setPosts(data.posts)
        })

         // ask back end for post
    }, [])

    const removeUser = () =>{
        // ask back end
        back()
    }

    function sendOrbit() {
        const path = `/sendfriendreq/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(props.userID)}`
        axios.post(`${serverpath}${path}`).then(res => {
            const data = res.data
            console.log(data)

            if (data.status === "no error") {
                setFriendStatus("viewer sent request")
            }
        })
    }

    function unorbit() {
        const path = `/removefriend/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(props.userID)}`
//        const path = `/sendfriendreq/3/1`
        axios.delete(`${serverpath}${path}`).then(res => {
            const data = res.data
            console.log(data)

            if (data.status === "no error") {
                setFriendStatus("no requests")
            }
        })
    }

    function rejectOrbitRequest() {
        const path = `/rejectfriendreq/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(props.userID)}`
//        const path = `/sendfriendreq/3/1`
        axios.delete(`${serverpath}${path}`).then(res => {
            const data = res.data
            console.log(data)

            if (data.status === "no error") {
                setFriendStatus("no requests")
            }
        })
    }

    function acceptFriendRequest() {
        const path = `/sendfriendreq/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(props.userID)}`
//        const path = `/sendfriendreq/3/1`
        axios.post(`${serverpath}${path}`).then(res => {
            const data = res.data
            console.log(data)

            if (data.status === "no error") {
                setFriendStatus("already friends")
            }
        })
    }

    const table = {
        "" : null,
        "already friends": <div onClick={unorbit}>Unorbit this buddy</div>,
        "no requests": <div onClick={sendOrbit}>Request orbit buddy</div>,
        "viewer sent request": <div>Orbit request pending</div>,
        "viewed person sent request": <div className="flex flex-col">
            <div onClick={acceptFriendRequest}>Accept orbit request</div>
            <div onClick={rejectOrbitRequest}>Reject orbit request</div>
        </div>
    }

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

                <button onClick={sendOrbit} className="flex flex-row cursor-pointer mb-3 hover:text-purple-300">
                    <img src="./addwhite.png" className="h-5 aspect-square translate-y-0.5 mr-2"/>
                    { table[friendStatus] }
                </button>

                {/* <button onClick={dm} className="flex flex-row cursor-pointer hover:text-purple-300">
                    <img src="./whitehole.png" className="h-5 aspect-square translate-y-0.5 mr-2"/>
                    <p>Launch wormhole chat</p>
                </button> */}

            </div>

            <div className="translate-x-1/3 w-2/3 pl-7 md:pl-36 lg:pl-0 lg:translate-x-0 lg:w-auto">
                <PlanetCanvas planet={user.planet}/>
            </div>

            </div>
            {
                posts? posts.map((post,index) => 
                (<Post key = {index} post = {post}/>)
                ) : null
            }

        </div>: null}

    </div>
        
    );
}

export default OtherProfile;