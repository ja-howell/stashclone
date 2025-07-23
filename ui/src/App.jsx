// import { useState } from 'react'
import { useEffect, useState } from 'react';
import './App.css'
import ItemCard from './components/ItemCard.jsx';

const API_URL = import.meta.env.VITE_API_URL;
const API_OPTIONS = {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json',
  },
};

const App = () => {
  const [items, setItems] = useState([]);
  const fetchItems = async () => {
    try {
      const response = await fetch(API_URL + "stashitems", API_OPTIONS);
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

      </div>
      <section className="all-items">
      <h2 className="text-xl font-bold mb-4">Your Stash!</h2>
      <div className="item-list grid grid-cols-1 xs:grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        {items.map((item, index) => (
          console.log(`Item ${index}:`, item),
          <ItemCard key={item.id} item={item} />
        ))}
      </div>
    </section>

    </main>

  )
}


export default App
