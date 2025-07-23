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
<div className="item-card border border-gray-300 rounded-md p-4 bg-white shadow hover:shadow-md transition">
  <h3 className="text-lg font-semibold mb-2">Name: {item.item_name}</h3>
  <p className="text-gray-600">Type: {item.item_type}</p>
</div>
  )
}

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

      </div>
      <section className="all-items">
      <h2 className="text-xl font-bold mb-4">Your Stash!</h2>
      <div className="item-list grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
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
