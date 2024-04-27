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
        <div style={{ display: 'flex',flexDirection: 'column', alignItems: 'center', justifyContent: 'center', width: '50%'}}>
            
            <input 
                type="text" 
                placeholder={type === "Start" ? "Insert your start wikipedia title here" : "Insert your end wikipedia title here"} 
                onChange={handleChange} 
                style={{ 
                    marginTop: '10px',
                    height:'30px',
                    borderRadius:'10px',
                    padding:'10px',
                    backgroundColor:'#200000',
                    border:'3px solid white',
                    color:'white',
                    textAlign: 'center',
                    width: '100%',
                    outline: 'none',
                }}
                    required 
            />
            {titleSuggestion.suggestions.length > 0 && (
                <ul style={{ margin: '0px', display: 'flex', flexDirection: 'column', paddingLeft: '0px', width: '100%'}}>
                    {titleSuggestion.suggestions.slice(0,5).map((title,index) => (
                        <button
                            key={index}
                            style={{
                                alignItems: 'center',
                                justifyContent: 'center',
                                display: 'flex',
                                padding: '10px',
                                textAlign: 'center',
                                cursor: 'pointer',
                                backgroundColor:'#200000',
                                color:'white',
                                border: 'none',
                            }}
                        >{title}</button>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default SearchInput;