import { React, useState } from "react";
import Planet from "./Planet.jsx";
import { Canvas } from "@react-three/fiber";

function MyProfile(props) {
    return (
        <Canvas className="cursor-pointer">
            <Planet></Planet>
        </Canvas>
    );
}

export default MyProfile;