import { Fragment, useEffect, useState } from "react"
import { useDebounce } from "use-debounce"
function SearchInput({type,setLinkValue}) {
    const [inputValue, setInputValue] = useState('');
    const [titleSuggestion,setTitleSuggestion] = useState({
        title: '',
        suggestions: [],
    })
    // const [valueInput, setValueInput] = useState(titleSuggestion.title)
    const [titleToFindSuggestion] = useDebounce(titleSuggestion.title,500)

    useEffect(() => {
        searchSuggestions(titleSuggestion.title)
    }, [titleToFindSuggestion])

    async function searchSuggestions(value){
        const realValue = value.replace(" ", "_");
        if (value !== '') {
            try {
                const response = await fetch(
                    `https://api.allorigins.win/get?url=${encodeURIComponent(
                        `https://en.wikipedia.org/w/api.php?action=query&format=json&list=search&srsearch=${realValue}`
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
            setInputValue(value)
            setTitleSuggestion({
                title: value,
                suggestions: [],
            })
        }
    }

    const handleSuggestionClick = (title) => {
        setLinkValue(title, type);
        setInputValue(title)
        setTitleSuggestion({
          title: title,
          suggestions: [],
        });
      };

    return (
        <div className="flex flex-col">
            <input 
                type="text" 
                placeholder={type}
                value = {inputValue}
                className="bg-[#8884bc] text-center w-full p-4 rounded-md border-black border-[3px] focus:outline-none placeholder-center placeholder-white text-[#f8ec34] text-2xl font-gingerCat shadowButton"
                onChange={handleChange}
                required
            />
            {titleSuggestion.suggestions.length > 0 && (
                <ul className="flex flex-col z-10 w-full mt-1 bg-[#8884bc] rounded-md border-black border-[3px]">
                  {titleSuggestion.suggestions.slice(0,5).map((title, index) => (
                    <li
                      key={index}
                      className="px-4 py-2 text-center cursor-pointer hover:bg-[#68649c] text-white text-xl font-gingerCat"
                      onClick={() => handleSuggestionClick(title)}
                    >
                      {title}
                    </li>
                  ))}
                </ul>
            )}
        </div>
  );
}

export default SearchInput;
