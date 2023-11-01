import { React, useState } from "react";

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
            </div>
        </div>
    );
}

function SearchUsers(props) {
    const samplePeople = [
        {username: "Vic", user_pic_url: "/assets/user0-pfp.jpg"}
    ]

    const [searchTerm, setSearchTerm] = useState("")
    const [people, setPeople] = useState(null)

    useEffect(() => {
        // ask back end
        setPeople(samplePeople)

    },[])
    return (
        <div>
            <input type="text" value={searchTerm} onChange={e => setSearchTerm(e.target.value)}>

            </input>
            {people.map(
                (person, index) => (
                    person.username.subString(searchTerm) ? 
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