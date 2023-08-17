import { TSayka } from "@/types/Sayka";
import { TUser } from "@/types/User";

import { TComment, TTag } from "./type";

// user
export const mock_user_1: TUser = {
  id: "hahaha-1",
  name: "hahaha",
  img: "/icon3.png",
  // favorites: 10,
  // sayka_count: 3,
};

export const mock_user_2: TUser = {
  id: "teyang-1",
  name: "teyang",
  img: "/icon2.jpg",
  // favorites: 100,
  // sayka_count: 30,
};

export const mock_users: TUser[] = [mock_user_1, mock_user_2];

// tag
export const mock_tag_1: TTag = "react";

export const mock_tag_2: TTag = "typescript";

export const mock_tag_3: TTag = "プロダクトを作りたい";

export const mock_tag_4: TTag = "designを作りたい";

export const mock_tags: TTag[] = [
  mock_tag_1,
  mock_tag_2,
  mock_tag_3,
  mock_tag_4,
];

// sayka
export const mock_saykas: TSayka[] = [
  {
    id: "1",
    title: "成果物共有アプリ『Sayka』",
    description:
      "このアプリは、成果物を多くの人に共有することで、より良い成果物を作ることを目的としたアプリです。",
    favorites: 13,
    comments: 3,
    github_url: "https://github.com/39TO",
    slide_url:
      "https://docs.google.com/presentation/d/1DsnKwxItsMPsA7-WeWLvDOIMLbApLPPFBtJD5WLemy4/edit?usp=sharing",
    app_url: "https://nextjs.org/docs",
    figma_url:
      "https://www.figma.com/file/1DsnKwxItsMPsA7-WeWLvDOIMLbApLPPFBtJD5WLemy4/edit",
    article_url: "https://qiita.com/39_to",
    user_id: mock_user_1.id,
    tags: [mock_tag_1, mock_tag_2, mock_tag_3, mock_tag_4],
    is_favorite: false,
    created_at: "2021-10-10",
  },
  {
    id: "2",
    title:
      "最強のアプリ-最強のアプリ-最強のアプリ-最強のアプリ-最強のアプリ-最強のアプリ-最強のアプリ-",
    description:
      "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
    favorites: 1,
    comments: 5,
    github_url: "https://github.com/39TO",
    slide_url:
      "https://docs.google.com/presentation/d/1DsnKwxItsMPsA7-WeWLvDOIMLbApLPPFBtJD5WLemy4/edit?usp=sharing",
    user_id: mock_user_1.id,
    tags: [mock_tag_1],
    is_favorite: true,
    created_at: "2021-10-10",
  },
  {
    id: "3",
    title: "Sayka 2",
    description: "console.log('Hello World')",
    favorites: 1,
    comments: 5,
    slide_url:
      "https://docs.google.com/presentation/d/1DsnKwxItsMPsA7-WeWLvDOIMLbApLPPFBtJD5WLemy4/edit?usp=sharing",
    app_url: "https://nextjs.org/docs",
    user_id: mock_user_2.id,
    tags: [],
    is_favorite: true,
    created_at: "2021-10-10",
  },
  {
    id: "4",
    title: "Sayka 2",
    description: "Sayka 2 description",
    favorites: 1,
    comments: 5,
    user_id: mock_user_2.id,
    tags: [mock_tag_1, mock_tag_2, mock_tag_3],
    is_favorite: false,
    created_at: "2021-10-10",
  },
];

export const mock_saykaNull: TSayka[] = [];

// saykaComments
export const mock_saykaComments_saykaId1: TComment[] = [
  {
    id: 1,
    contents:
      "面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！",
    created_at: "2021-10-10",
    user: mock_user_2,
  },
  {
    id: 2,
    contents:
      " > console.log(data) \n\nこれは一人で開発されましたか？よければ一緒にやりたいです！ \n\n```console.log('Hello World')```ga`aa`",
    created_at: "2021-10-10",
    user: mock_user_1,
  },
  {
    id: 3,
    contents:
      "これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！",
    created_at: "2021-10-10",
    user: mock_user_2,
  },
];
