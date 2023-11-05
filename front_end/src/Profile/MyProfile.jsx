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

function MyProfile(props) {
    const toggleFriendsList = props.toggleFriendsList
    const toggleSettings = props.toggleSettings
    const toggleHomepage = props.toggleHomepage

    const duppy = {
        username: "Duppy", 
        homePlanet: "Earth",
        pfp: "./ayylmao.webp",
        bio: "My unmatched perspicacity coupled with sheer indefitigability makes me a feared opponent in any realm of human endeavor.\n https://okhan.me"
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
        <div className="flex flex-col lg:flex-row justify-center align-middle">

    
            <div className="flex flex-col ml-5 mr-5 lg:ml-0 lg:mr-0 lg:w-80">

                <div className="flex flex-row mb-4">

                    <img src={user.pfp} alt="My Profile Picture" className="w-20 aspect-square rounded-full"/>

                    <div className="flex flex-col ml-4">
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

        </div>: null}

    </div>
        
    );
}

export default MyProfile;