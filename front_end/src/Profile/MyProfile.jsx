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

function Post() {
    return(
        <div className="flex flex-col w-5/12 mx-auto mb-5 justify-center align-middle bg-slate-300 bg-opacity-90 text-black rounded-lg">

            <div className="flex flex-row p-3 justify-between">
                <div className="flex flex-row justify-center align-middle ">
                    <img src="./ayylmao.webp" alt="Profile Picture" className="w-10 aspect-square rounded-full"/>
                    <p className="ml-2 mt-2">Duppy</p>
                </div>
                <p className="mr-2 mt-2">1h</p>
            </div>



            <p className="mt-2 pl-5"> Finally leaving this planet lmao 😂</p>

            <img src="./swag.jpg" className="m-4" alt="" />


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
    
    useEffect(() => {
        // ask bak end
         setUser(duppy)
    }, [])

    return (
        <div className="flex flex-col">

        <button className="mb-5 w-fit ml-6 text-xl" onClick={toggleHomepage}> {'<--'} </button>

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

            <Post/>
            <Post/>
            <Post/>



        </div>: null}


    </div>
        
    );
}

export default MyProfile;