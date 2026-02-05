import React, { useState } from 'react';

const MultilineInput = () => {
  const [textAreaValue, setTextAreaValue] = useState('');

  const handleChange = (event) => {
    setTextAreaValue(event.target.value);
  };

  return (
    <div className="flex flex-col gap-4 p-4 w-full max-w-md">
      <textarea
        id="multiline-text"
        value={textAreaValue}
        onChange={handleChange}
        className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        rows={5}
        placeholder="Enter multiple IDs"
      />
      <button
        type="submit"
        className="self-end px-6 py-2 bg-blue-400 text-white font-semibold rounded-lg hover:bg-blue-650 transition-colors"
      >
        Search
      </button>
    </div>
  );
};

export default MultilineInput;
