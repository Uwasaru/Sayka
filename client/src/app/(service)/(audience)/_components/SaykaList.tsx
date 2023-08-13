"use client";

import { useState } from "react";

import { Sayka, TSayka } from "./Sayka";
import { TagFilter } from "./TagFilter";

// TODO: APIから取得したデータを表示する
export const SaykaMockData: TSayka[] = [
  {
    id: 1,
    title: "成果物共有アプリ『Sayka』",
    description:
      "このアプリは、成果物を多くの人に共有することで、より良い成果物を作ることを目的としたアプリです。",
    like_count: 13,
    comment_count: 3,
    github_url: "https://github.com/39TO",
    slide_url:
      "https://docs.google.com/presentation/d/1DsnKwxItsMPsA7-WeWLvDOIMLbApLPPFBtJD5WLemy4/edit?usp=sharing",
    application_url: "https://nextjs.org/docs",
    user: {
      id: 1,
      name: "39TO",
      icon: "/icon.jpg",
    },
    tags: [
      {
        id: 1,
        name: "react",
      },
      {
        id: 2,
        name: "nextjs",
      },
    ],
  },
  {
    id: 2,
    title:
      "最強のアプリ-最強のアプリ-最強のアプリ-最強のアプリ-最強のアプリ-最強のアプリ-最強のアプリ-",
    description:
      "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
    like_count: 1,
    comment_count: 5,
    github_url: "https://github.com/39TO",
    slide_url:
      "https://docs.google.com/presentation/d/1DsnKwxItsMPsA7-WeWLvDOIMLbApLPPFBtJD5WLemy4/edit?usp=sharing",
    application_url: "https://nextjs.org/docs",
    user: {
      id: 1,
      name: "39TO",
      icon: "/icon.jpg",
    },
    tags: [
      {
        id: 1,
        name: "java",
      },
      {
        id: 2,
        name: "サッカー選手",
      },
    ],
  },
  {
    id: 3,
    title: "Sayka 2",
    description: "console.log('Hello World')",
    like_count: 1,
    comment_count: 5,
    github_url: "https://github.com/39TO",
    slide_url:
      "https://docs.google.com/presentation/d/1DsnKwxItsMPsA7-WeWLvDOIMLbApLPPFBtJD5WLemy4/edit?usp=sharing",
    application_url: "https://nextjs.org/docs",
    user: {
      id: 1,
      name: "39TO",
      icon: "/icon.jpg",
    },
    tags: [
      {
        id: 1,
        name: "java",
      },
      {
        id: 2,
        name: "起業家",
      },
    ],
  },
  {
    id: 4,
    title: "Sayka 2",
    description: "Sayka 2 description",
    like_count: 1,
    comment_count: 5,
    github_url: "https://github.com/39TO",
    slide_url:
      "https://docs.google.com/presentation/d/1DsnKwxItsMPsA7-WeWLvDOIMLbApLPPFBtJD5WLemy4/edit?usp=sharing",
    application_url: "https://nextjs.org/docs",
    user: {
      id: 1,
      name: "39TO",
      icon: "/icon.jpg",
    },
    tags: [
      {
        id: 1,
        name: "java",
      },
      {
        id: 2,
        name: "起業家",
      },
    ],
  },
];

export const SaykaList = () => {
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
      <div className="mt-5 space-y-5">
        {SaykaMockData.map((sayka) => (
          <Sayka key={sayka.id} data={sayka} changeFilterTag={clickTag} />
        ))}
      </div>
    </div>
  );
};
