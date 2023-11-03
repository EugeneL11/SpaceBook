import { React, useState, useEffect } from "react";
import {userID} from "../Static.js"
import Planet from "./Planet.jsx";
import { Canvas } from "@react-three/fiber";

function MyProfile(props) {
    const toggleFriendsList = props.toggleFriendsList
    const toggleSettings = props.toggleSettings
    const toggleHomepage = props.toggleHomepage

    const duppy = {username: "Duppy"}
    const [user, setUser] = useState(null)
    
    useEffect(() => {
        // ask bak end
         setUser(duppy)
    }, [])
    
    return (
        <div className="flex flex-col">
        <h1>This is the My Profile component</h1>
        <button onClick={toggleSettings}>Go to Settings</button>
        <button onClick={toggleFriendsList}>Go to Friends List</button>
        <button onClick={toggleHomepage}>Go to Homepage</button>
        { user ? 
        <div className="flex flex-row justify-center align-middle">
            <p>{user.username}</p>
            <Canvas className="cursor-pointer">
                <Planet></Planet>
            </Canvas>
        </div>: null}
    
    </div>
        
    );
}

export default MyProfile;