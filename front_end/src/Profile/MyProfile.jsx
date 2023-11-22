import { React, useState, useEffect } from "react";
import currentUser from "../Static.js"
import Planet from "./Planet.jsx";
import { Canvas } from "@react-three/fiber";
import axios from 'axios'
function PlanetCanvas () {
    return (
    <Canvas className="cursor-pointer mt-5 md:mt-0">
        <Planet planet="pluto"/>
    </Canvas>
    )
}

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
        <div className="flex flex-col w-11/12 lg:w-5/12 mx-auto mb-10 justify-center align-middle bg-slate-300 bg-opacity-90 text-black rounded-lg">
            <div className="relative w-100 h-100">
            {imageNum > 0 ? 
                // <div className="absolute text-purple-500 pb-2 pr-2 pl-2 bg-slate-300 bg-opacity-60 rounded-full text-7xl top-52 z-40 cursor-pointer hover:text-purple-400" onClick={togglePrevImage}> {"←"} </div> : null
                <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer translate-x-10 translate-y-10 rotate-180" onClick={togglePrevImage} /> : null
            }
            {imageNum < imageCount - 1 ? 
                // <div className="absolute text-purple-500 pb-2 pr-2 pl-2 bg-slate-300 bg-opacity-60 rounded-full text-7xl cursor-pointer left-85-percent top-52 z-40 hover:text-purple-400" onClick={toggleNextImage}> {"→"} </div> : null
                <img src="./ar.png" className="absolute w-10 p-2 bg-slate-300 bg-opacity-80 rounded-full text-7xl top-52 z-40 cursor-pointer right-0 -translate-x-10 translate-y-10"  onClick={toggleNextImage}/> : null

            }
                </div>
            <div className="flex flex-row p-3 justify-between">
                <div className="flex flex-row justify-center align-middle pl-2 pt-2 ">
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

function MyProfile(props) {
    const toggleFriendsList = props.toggleFriendsList
    const toggleSettings = props.toggleSettings
    const toggleHomepage = props.toggleHomepage

    const duppy = {
        username: "Duppy", 
        homePlanet: "Earth",
        pfp: "./ayylmao.webp",
        bio: "My unmatched perspicacity coupled with sheer indefitigability makes me a feared opponent in any planet of this solar system.\n https://okhan.me/"
    }
    const [user, setUser] = useState(null)
    const [posts, setPosts] = useState(null)
    const examplePost = {
        username: "Duppy",
        pfp : "./ayylmao.webp",
        caption: "Finally leaving this planet lmao 😂",
        date: "Nov 7th",
        images: ["./swag.jpg", "./ayylmao.webp"],
        videos: [],
    }
    const examplePosts =[ examplePost,examplePost]
    useEffect(() => {
        // ask bak end for user
         setUser(duppy)
         // ask back end for post
         setPosts(examplePosts);
    }, [])

    return (
        <div className="flex flex-col">
        
        <button className="mb-5 w-fit ml-6 text-3xl hover:text-purple-300" onClick={toggleHomepage}> {'←'} </button>

        { user ? 
        
        <div className="flex flex-col justify-center align-middle">

        <div className="flex flex-col lg:flex-row justify-center align-middle mb-14">
    
            <div className="flex flex-col ml-5 mr-5 lg:ml-0 lg:mr-0 lg:w-80">

                <div className="flex flex-row mb-4 align-middle">

                    <img src={user.pfp} alt="My Profile Picture" className="w-20 aspect-square rounded-full"/>

                    <div className="flex flex-col ml-4 justify-center align-middle">
                        <p className="text-xl">{user.username}</p>
                        <p>from {user.homePlanet}</p>
                    </div>

                </div>

                <p className="mb-4">{user.bio}</p>

                <button onClick={toggleSettings} className="flex flex-row cursor-pointer mb-3 hover:text-purple-300">
                    <img src="./gear.png" className="h-5 aspect-square translate-y-0.5 mr-2"/>
                    <p>Edit profile settings</p>
                </button>

                <button onClick={toggleFriendsList} className="flex flex-row cursor-pointer hover:text-purple-300">
                    <img src="./sun.png" className="h-5 aspect-square translate-y-0.5 mr-2"/>
                    <p>View orbit buddies</p>
                </button>

            </div>



            <div>
                <PlanetCanvas/>
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

export default MyProfile;
