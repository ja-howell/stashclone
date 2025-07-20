// import { useState } from 'react'
import { useEffect, useState } from 'react';
import './App.css'

const API_URL = 'http://localhost:8080/stashitems';
const API_OPTIONS = {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json',
  },
};

const ItemCard = ({item}) => {
  console.log("Rendering ItemCard for", item.name);
  return (
    <div className="item-card">
      <h3>Name: {item.item_name}</h3>
      <p>Type: {item.item_type}</p>
    </div>
  )
}

// const ItemCard = (props) => {
//   console.log('ItemCard props:', props);
//   const { item } = props;

//   return (
//     <div className="item-card">
//       <h3>{item.name}</h3>
//       <p>{item.type}</p>
//     </div>
//   );
// };


const App = () => {
  const [items, setItems] = useState([]);
  const fetchItems = async () => {
    try {
      const response = await fetch(API_URL, API_OPTIONS);
      const data = await response.json();
      console.log('Response:', response);
      console.log('Data:', data);

      setItems(data || []);
      } catch (error) {
      console.error('Fetch error:', error);
    }
  }

  useEffect(() => {
    fetchItems();
  }, []);
  useEffect(() => {
  console.log('Items updated:', items.length);
}, [items]);


  console.log("Rendering", items.length, "ItemCards");

  return (
    <main>
      <div className="header">
        <header>
        <h1>Stash Clone</h1>
        </header>

        <section className="items">
        <h2>Your Stash!</h2>
          <div className="items-list">
            {items.map((item, index) => (
              console.log(`Item ${index}:`, item),
              <ItemCard key={item.id} item={item} />
            ))}
          </div>
        </section>
      </div>
    </main>

  )
}


export default App
