
import {useState} from "react";


export default function Form(){
    const [exercise, setExercise] = useState({name: "", times: "", weight: ""});

    function handleSubmit(e){
      e.preventDefault();
      console.log(exercise);
    }

    return (
        <>
          <form>
              <label>
                Exercício: <input name="name" onChange={(e) => setExercise({...exercise, name: e.target.value})} />
              </label><br></br>
    
              <label>
                Repetições: <input name="times" onChange={(e) => setExercise({...exercise, times: e.target.value})} />
              </label><br></br>
    
              <label>
                Peso: <input name="weight" onChange={(e) => setExercise({...exercise, weight: e.target.value})}/> kg
              </label><br></br>
    
              <button type="submit" onClick={(e) => handleSubmit(e)}>Enviar</button>
          </form>
        </>
      )
}