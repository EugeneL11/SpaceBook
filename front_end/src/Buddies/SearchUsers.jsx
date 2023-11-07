import { React, useState , useEffect } from "react";
import { userID } from "../Static.js";

function Person(props) {
    const user_pic_url = props.user_pic_url
    return (
        <div>
            <div class="flex-row">
                <img 
                    src={props.user_pic_url}
                    alt={props.username} 
                    class="rounded-lg"
                ></img>
                <p>{props.username}</p>
                {/* <button onClick={() => {toggleOtherProfile(props.username, toggleSearchUser)}}> See Other Profile: {person}</button> */}
            </div>
        </div>
    );
}

function SearchUsers(props) {
    const samplePeople = [
        {username: "Vic", user_pic_url: "/assets/user0-pfp.jpg"},
        {username: "Kevin", user_pic_url: "/assets/user0-pfp.jpg"}
    ]

    const toggleHomepage = props.toggleHomepage
    const toggleOtherProfile = props.toggleOtherProfile
    const toggleSearchUser = props.toggleSearchUser
    const exampleFriends = ["Kevin", "Omar" , "Raine", "Eugene"]

    const [searchTerm, setSearchTerm] = useState("")
    const [people, setPeople] = useState(null)

    useEffect(() => {
        // ask back end
        setPeople(samplePeople)

    },[])
    return (
        <div className="flex flex-col">
            <h1>This is the Search Users component</h1>
            <button onClick={toggleHomepage}>Go to Homepage Screen</button>
            <input type="text" value={searchTerm} onChange={e => setSearchTerm(e.target.value)}>

            </input>
            {samplePeople.map(
                (person, index) => (
                    person.username.substring(searchTerm) ? 
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