import { React, useState } from "react";
import backPic from '../images/back.png';

function Settings(props) {
    const toggleLogin = props.toggleLogin
    const toggleMyProfile = props.toggleMyProfile
    return (
    <div className="flex flex-col items-center">
        <div className=" flex flex-col bg-white text-black text-start text-lg m-10 md:py-6 sm:px-16 lg:px-24 p-6 rounded-xl w-3/4 md:w-1/2 min-w-fit">
            <div className="flex justify-between items-center">
                <img onClick={toggleMyProfile} src={backPic} alt="a back arrow button" className="w-10 cursor-pointer"></img>
                <h1 className="text-center text-3xl py-4">Settings</h1>
                <img onClick={toggleMyProfile} src={backPic} alt="a back arrow button" className="w-10 invisible"></img>
            </div>

            <label className="mt-6">Change Username: </label>
            <input className="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" placeholder="Enter New Username" type="text"></input>
        
            <div className="mt-6">Change Profile Picture: </div>
            <input type="file" className="form-input text-sm"></input>

            <div className="mt-6">Edit Bio:</div>
            <textarea className="form-textarea border-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" rows="3" placeholder=" Edit your bio..."></textarea>

            <div className="mt-6">Change Home Planet: </div>
            <select className="form-select border-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0">
                <option value="mercury">mercury</option>
                <option value="venus">venus</option>
                <option value="earth">earth</option>
                <option value="mars">mars</option>
                <option value="jupiter">jupiter</option>
                <option value="saturn">saturn</option>
                <option value="uranus">uranus</option>
                <option value="neptune">neptune</option>
                <option value=""></option>
                <option value=""></option>
            </select>
            
            <button classNameName="bg-purple-300 px-7 py-3 mt-6 w-fit self-center rounded-lg">Apply Changes</button>

            <button classNameName="mt-10" onClick={toggleLogin}>Log Out</button>
            <button>Delete Account</button>
        </div>
    </div>
    );

}

export default Settings;