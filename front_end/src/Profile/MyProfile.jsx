import { React, useState, useEffect } from "react";
import {userID} from "../Static.js"
import Planet from "./Planet.jsx";
import { Canvas } from "@react-three/fiber";

function PlanetCanvas () {
    return (
    <Canvas className="cursor-pointer">
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
        pfp: "./ayylmao.webp"
    }
    const [user, setUser] = useState(null)
    
    useEffect(() => {
        // ask bak end
         setUser(duppy)
    }, [])

    return (
        <div className="flex flex-col">

        <button className="mb-5" onClick={toggleHomepage}>Go to Homepage</button>


        { user ? 
        <div className="flex flex-row justify-center align-middle">

            <div className="flex flex-col">

                

                <div className="flex flex-row align-middle justify-center">

                    <img src={user.pfp} alt="My Profile Picture" className="w-20 aspect-square rounded-full"/>

                    <div className="flex flex-col ml-4">
                        <p>{user.username}</p>
                        <p>from {user.homePlanet}</p>
                    </div>

                </div>

                <button onClick={toggleSettings}>Go to Settings</button>
                <button onClick={toggleFriendsList}>Go to Friends List</button>

            </div>



            <div>
            <PlanetCanvas/>
        </div>

        </div>: null}

    </div>
        
    );
}

export default MyProfile;