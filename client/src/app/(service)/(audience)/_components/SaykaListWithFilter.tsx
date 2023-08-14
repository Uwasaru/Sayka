"use client";

import { useState } from "react";

import { mock_saykas } from "@/api";

import { SaykaList } from "./SaykaList";
import { TagFilter } from "./TagFilter";

export const SaykaListWithFilter = () => {
  const [tagState, setTagState] = useState<string>("ALL");

  const changeFilterTag = (tag: string) => {
    setTagState(tag);
  };

  const handleSubmit = (tag: string) => {
    console.log("Submitted:", tag);
  };

  const clickTag = (tag: string) => {
    changeFilterTag(tag);
    if (tagState !== tag) {
      handleSubmit(tag);
    }
  };

  return (
    <div>
      <TagFilter
        tag={tagState}
        changeFilterTag={changeFilterTag}
        onSubmit={handleSubmit}
      />
      <SaykaList sayka={mock_saykas} changeFilterTag={clickTag} />
    </div>
  );
};