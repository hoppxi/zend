import React, { useRef, useState } from "react";

export const SearchBar: React.FC = () => {
  const [q, setQ] = useState("");
  const inputRef = useRef<HTMLInputElement | null>(null);

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const query = q.trim();
    if (!query) return;
    const url = `https://www.google.com/search?q=${encodeURIComponent(query)}`;
    window.location.href = url;
  };

  return (
    <form className={`search ${q ? "active" : ""}`} onSubmit={onSubmit}>
      <input
        ref={inputRef}
        className="search-input"
        placeholder="Search the web…"
        value={q}
        onChange={(e) => setQ(e.target.value)}
        onFocus={() => document.documentElement.classList.add("search-focus")}
        onBlur={() => document.documentElement.classList.remove("search-focus")}
      />
      <button className="search-btn" aria-label="Search">
        ↵
      </button>
    </form>
  );
};
