import { React, useState, useEffect } from "react";
import {userID} from "../Static.js"
import Planet from "./Planet.jsx";
import { Canvas } from "@react-three/fiber";

function PlanetCanvas () {
    return (
    <Canvas className="cursor-pointer mt-5 md:mt-0"
        // camera={{ position: [0, 0, 1], fov: 50 }}
        >
        <Planet/>
    </Canvas>
    )
}

function Post(props) {
    const imageCount = props.post.images.length;
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
        <div className="flex flex-col w-5/12 mx-auto mb-5 justify-center align-middle bg-slate-300 bg-opacity-90 text-black rounded-lg">
            <div className="relative w-100 h-100">
            {imageNum > 0 ? 
                <div className="absolute text-green-300 text-9xl top-52 z-50 hover:text-green-400" onClick={togglePrevImage}> {"<"} </div> : null}
            {imageNum < imageCount - 1 ? 
                <div className="absolute text-green-300 text-9xl left-85-percent top-52 z-50 hover:text-green-400" onClick={toggleNextImage}> {">"} </div> : null}
                </div>
            <div className="flex flex-row p-3 justify-between">
                <div className="flex flex-row justify-center align-middle ">
                    <img src={props.post.pfp} alt="Profile Picture" className="w-10 aspect-square rounded-full"/>
                    <p className="ml-2 mt-2">{props.post.username}</p>
                </div>
                <p className="mr-2 mt-2">{props.post.date}</p>
            </div>



            <p className="mt-2 pl-5"> {props.post.caption} 😂</p>

            <img src={props.post.images[imageNum]} className="m-4" alt="" />


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
        bio: "My unmatched perspicacity coupled with sheer indefitigability makes me a feared opponent in any planet in this solar system.\n https://okhan.me"
    }
    const [user, setUser] = useState(null)
    const [posts, setPosts] = useState(null)
    const examplePost = {
        username: "duppy",
        pfp : "./ayylmao.webp",
        caption: "Finally Leaving this planet lmao",
        date: "Nov 7th",
        images: ["./ayylmao.webp", "./swag.jpg"],
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

        <button className="mb-5 w-fit ml-6 text-3xl" onClick={toggleHomepage}> {'<--'} </button>

        { user ? 
        
        <div className="flex flex-col justify-center align-middle">

        <div className="flex flex-col lg:flex-row justify-center align-middle mb-14">
    
            <div className="flex flex-col ml-5 mr-5 lg:ml-0 lg:mr-0 lg:w-80">

                <div className="flex flex-row mb-4 align-middle">

                    <img src={user.pfp} alt="My Profile Picture" className="w-20 aspect-square rounded-full"/>

                    <div className="flex flex-col ml-4 justify-center align-middle">
                        <p>{user.username}</p>
                        <p>from {user.homePlanet}</p>
                    </div>

                </div>

                <p className="mb-4">{user.bio}</p>

                <button onClick={toggleSettings} className="flex flex-row cursor-pointer ">
                    <img src="./gear.png" className="h-5 aspect-square translate-y-0.5 mr-2"/>
                    <p>Edit profile settings</p>
                </button>

                <button onClick={toggleFriendsList} className="flex flex-row cursor-pointer ">
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