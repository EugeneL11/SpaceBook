import { React, useState } from "react";
import backPic from '../images/back.png';

function Settings(props) {
    const toggleLogin = props.toggleLogin
    const toggleMyProfile = props.toggleMyProfile
    return (
    <div className="flex flex-col items-center">
        <button className="w-fit ml-2 mr-auto text-5xl hover:text-purple-300" onClick={toggleMyProfile}>{'←'}</button>
        <div className=" flex flex-col bg-white text-black text-start text-lg md:py-6 sm:px-16 lg:px-24 mt-[-25px] p-6 rounded-xl w-3/4 md:w-1/2 min-w-fit">
            <h1 className="text-center text-3xl pt-2">Settings</h1>
            <label className="mt-4">Change Username: </label>
            <input className="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" placeholder="Enter New Username" type="text"></input>
        
            <div className="mt-4">Change Profile Picture: </div>
            <input type="file" className="form-input text-sm"></input>

            <div className="mt-4">Edit Bio:</div>
            <textarea className="form-textarea border-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" rows="3" placeholder=" Edit your bio..."></textarea>

            <div className="mt-4">Change Home Planet: </div>
            <select className="form-select border-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0">
                <option value="mercury">mercury</option>
                <option value="venus">venus</option>
                <option value="earth">earth</option>
                <option value="mars">mars</option>
                <option value="jupiter">jupiter</option>
                <option value="saturn">saturn</option>
                <option value="uranus">uranus</option>
                <option value="neptune">neptune</option>
                <option value="pluto">pluto</option>
            </select>
            
            <button className="bg-purple-300 hover:bg-purple-400 px-4 py-2 mt-5 w-fit self-center rounded-lg">Apply Changes</button>

            <button className="w-fit self-center mt-6 hover:text-blue-300" onClick={toggleLogin}>Log Out</button>
            <button className="w-fit self-center hover:text-red-600">Delete Account</button>
        </div>
    </div>
    );

}

export default Settings;