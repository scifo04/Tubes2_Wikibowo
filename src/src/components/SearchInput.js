import { json } from "d3";
import React, { useEffect, useState } from "react";
import { useDebounce } from "use-debounce";



function SearchInput({setLinkValue,type}) {
    const [titleSuggestion,setTitleSuggestion] = useState({
        title: '',
        suggestions: [],
    })
    const [titleToFindSuggestion] = useDebounce(titleSuggestion.title,300)

    useEffect(() => {
        searchSuggestions(titleSuggestion.title)
    }, [titleToFindSuggestion])

    async function searchSuggestions(value){
        if (value !== '') {
            try {
                const response = await fetch(
                    `https://api.allorigins.win/get?url=${encodeURIComponent(
                        `https://en.wikipedia.org/w/api.php?action=query&format=json&list=search&srsearch=${value}`
                    )}`
                );
                const data = await response.json()
                const jsonData = JSON.parse(data.contents);
                // console.log(jsonData.query.search.map(item => item.title))
                setTitleSuggestion({
                    title : value,
                    suggestions : jsonData.query.search.map(item => item.title)
                })
                
            } catch (error) {
                console.error('Error fetching suggestions:', error);
                setTitleSuggestion({
                    title: '',
                    suggestions: [],
                });
            }
        }
    }
        
        const handleChange = async (e) => {
            if (e !== null && e !== undefined && e !== '') {
                const value = e.target.value;
                setLinkValue(value,type);
                setTitleSuggestion({
                    title: value,
                    suggestions: [],
                })
        }
    }
    return (
        <div>
            <input 
                type="text" 
                placeholder={type === "Start" ? "Insert your start wikipedia title here" : "Insert your end wikipedia title here"}
                className="inoneline" 
                onChange={handleChange} 
                style={{ 
                    margin: '10px',
                    width:'500px',
                    height:'30px',
                    borderRadius:'20px',
                    padding:'10px',
                    backgroundColor:'#200000',
                    border:'3px solid white',
                    color:'white' }}
                    required 
            />
            <ul>
                {titleSuggestion.suggestions.map((title,index) => (
                    <li key={index}>{title}</li>
                ))}
            </ul>
        </div>
    );
};

export default SearchInput;