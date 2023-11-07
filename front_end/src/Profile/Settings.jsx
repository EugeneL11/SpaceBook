import { React, useState } from "react";

function Settings(props) {
    const toggleLogin = props.toggleLogin
    const toggleMyProfile = props.toggleMyProfile
    return (
    <div className="flex flex-col items-center">
        <div class=" flex flex-col bg-white text-black text-start text-lg m-10 md:py-6 sm:px-16 lg:px-24 p-6 rounded-xl w-3/4 md:w-1/2 min-w-fit">
            <h1 class="text-center text-3xl py-4">Settings</h1>

            <button class="w-fit self-center" onClick={toggleMyProfile}>Back</button>

            <label class="">Change Username: </label>
            <input class="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" placeholder="Enter New Username" type="text"></input>
        
            <div class="mt-6">Change Profile Picture: </div>
            <input type="file" class="form-input text-sm"></input>

            <div class="mt-6">Edit Bio:</div>
            <textarea class="form-textarea border-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" rows="3" placeholder=" Edit your bio..."></textarea>

            <div class="mt-6">Change Home Planet: </div>
            <select class="form-select border-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0">
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
            
            <button class="bg-purple-300 px-7 py-3 mt-6 w-fit self-center rounded-lg">Apply Changes</button>

            <button class="mt-10" onClick={toggleLogin}>Log Out</button>
            <button>Delete Account</button>
        </div>
    </div>
    );

}

export default Settings;