"use client";

import { FC } from "react";
import { BsFilterLeft } from "react-icons/bs";

type TProps = {
  tag: string;
  changeFilterTag: (tag: string) => void;
  onSubmit: (tag: string) => void;
};

export const TagFilter: FC<TProps> = ({ tag, changeFilterTag, onSubmit }) => {
  const handleSubmit = () => {
    if (!tag.trim()) {
      changeFilterTag("ALL");
    }
    onSubmit(tag);
  };

  return (
    <div className="flex items-center space-x-2">
      <BsFilterLeft size={24} />
      <div className="inline-block cursor-pointer rounded-2xl border border-gray-300 bg-gradient-to-r from-gray-400 to-gray-500 px-3 py-1 font-semibold shadow-md focus-within:ring-2 focus-within:ring-teal-400 hover:bg-gradient-to-r hover:from-gray-200 hover:to-gray-300">
        #{" "}
        <input
          type="tag"
          value={tag}
          onChange={(e) => changeFilterTag(e.target.value)}
          size={tag?.length || 1}
          className="bg-transparent focus:outline-none"
          onBlur={handleSubmit}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              e.preventDefault();
              handleSubmit();
            }
          }}
        />
      </div>
    </div>
  );
};
