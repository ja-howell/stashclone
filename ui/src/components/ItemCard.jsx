import React from "react";

const ItemCard = ({ item }) => {
  console.log("Rendering ItemCard for", item.name);
  return (
    <div className="item-card border border-gray-300 rounded-md p-4 bg-white shadow hover:shadow-md transition">
      <h3 className="text-lg font-semibold mb-2">Name: {item.item_name}</h3>
      <p className="text-gray-600">Type: {item.item_type}</p>
    </div>
  );
}

export default ItemCard;