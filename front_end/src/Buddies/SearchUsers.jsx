import { React, useState , useEffect } from "react";
import { userID } from "../Static.js";

function Person(props) {
    const user_pic_url = props.user_pic_url
    return (
        <div className="flex items-center w-1/2 bg-blue-500 space-x-4 rounded-md hover:cursor-pointer hover:bg-blue-300">
            <img 
                src={props.user_pic_url}
                alt={props.username} 
                className="w-12 h-12 rounded-full aspect-square p-2"
            ></img>
            <p className="text-lg">{props.username}</p>
            {/* <button onClick={() => {toggleOtherProfile(props.username, toggleSearchUser)}}> See Other Profile: {person}</button> */}
        </div>
    );
}

function SearchUsers(props) {
    const samplePeople = [
        {username: "Vic", user_pic_url: "./jupiter.jpg"},
        {username: "Kevin", user_pic_url: "./jupiter.jpg"}
    ]

    const toggleHomepage = props.toggleHomepage
    const toggleOtherProfile = props.toggleOtherProfile
    const toggleSearchUser = props.toggleSearchUser

    const [searchTerm, setSearchTerm] = useState("")
    const [people, setPeople] = useState(null)

    useEffect(() => {
        // ask back end for top 10
        setPeople(samplePeople)

    },[])

    return (
        <div className="flex flex-col justify-start items-center space-y-4">
            <div className="flex flex-start w-full">
                <button className="mb-5 w-fit ml-6 text-3xl hover:text-purple-300" onClick={toggleHomepage}> {'←'} </button>
            </div>
            <div className="flex w-1/2">
                <input 
                    type="text" 
                    value={searchTerm} 
                    onChange={e => setSearchTerm(e.target.value)}
                    className="w-full p-2 rounded-bl-md rounded-tl-md text-black"
                ></input>
                <div className="relative inset-y-0 right-0 flex items-center px-3 bg-white rounded-tr-md rounded-br-md">
                    <img
                        src="./search.png"
                        alt="search users"
                        className="w-6 h-6"
                    ></img>
                </div>
            </div>
            {samplePeople.map(
                (person, index) => (
                    person.username.toLowerCase().includes(searchTerm.toLowerCase()) ? 
                    <Person
                        key={index}
                        username={person.username} 
                        user_pic_url={person.user_pic_url}
                    ></Person> : null
                )
            )}
        </div>
    );

}

export default SearchUsers;