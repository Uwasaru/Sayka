"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { FC, useState } from "react";
import { useForm, SubmitHandler, FormProvider } from "react-hook-form";
import { AiFillGithub } from "react-icons/ai";
import { GrClose } from "react-icons/gr";
import { MdOutlineArticle } from "react-icons/md";
import { PiFigmaLogoDuotone } from "react-icons/pi";
import { TfiLayoutSlider } from "react-icons/tfi";
import { TfiWorld } from "react-icons/tfi";
import { z } from "zod";

import { Tag } from "@/ui/Tag";
import { ContentSubTitle, Explanation } from "@/ui/Text";

import { InputBlock } from "./InputBlock";
import { createSayka } from "@/api/sayka";
import { TUser } from "@/types/User";
import { redirect } from "next/navigation";
import { TSayka } from "@/types/Sayka";

import { useRouter } from "next/navigation";

const schema = z.object({
  title: z
    .string()
    .nonempty({ message: "必須項目です。" })
    .max(20, "20字以内で入力してください。"),
  description: z
    .string()
    .nonempty({ message: "必須項目です。" })
    .max(100, "100字以内で入力してください。"),
  github_url: z.union([
    z.literal(""),
    z.string().regex(/^https:\/\//, "URLを正しい形で入力してください。"),
  ]),
  slide_url: z.union([
    z.literal(""),
    z.string().regex(/^https:\/\//, "URLを正しい形で入力してください。"),
  ]),
  app_url: z.union([
    z.literal(""),
    z.string().regex(/^https:\/\//, "URLを正しい形で入力してください。"),
  ]),
});

type FormData = {
  title: TSayka["title"];
  description: TSayka["description"];
  github_url?: TSayka["github_url"];
  figma_url?: TSayka["figma_url"];
  slide_url?: TSayka["slide_url"];
  article_url?: TSayka["article_url"];
  app_url?: TSayka["app_url"];
};

type TProps = {
  user: TUser;
  token: string;
};

export const SaykaForm: FC<TProps> = ({ user, token }) => {
  const [tags, setTags] = useState<string[]>([]);
  const [tagError, setTagError] = useState<string | null>(null);
  const methods = useForm<FormData>({
    resolver: zodResolver(schema),
  });
  const router = useRouter();

  const onSubmit: SubmitHandler<FormData> = (data) => {
    const submitData = {
      ...data,
      tags,
    };
    createSayka(
      {
        user_id: user.id,
        ...submitData,
      },
      token
    ).then((res) => {
      if (res.type === "error") return;
      console.log(res.value.data.id);
      router.push(`/shareSaykaInformation/${res.value.data.id}`);
    });
  };

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value;
    const lastChar = value[value.length - 1];

    if (value.length > 12) {
      setTagError("タグは12文字以下である必要があります。");
      return;
    } else {
      setTagError(null);
    }

    if (lastChar === "," && tags.length < 5) {
      setTags([...tags, value.slice(0, -1)]);
      event.target.value = "";
    }
  };

  const handleTagRemove = (index: number) => {
    const newTags = [...tags];
    newTags.splice(index, 1);
    setTags(newTags);
  };

  const preventFormSubmitOnEnter = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" && e.currentTarget === e.target) {
      e.preventDefault();
    }
  };

  return (
    <FormProvider {...methods}>
      <form onKeyDown={preventFormSubmitOnEnter}>
        <InputBlock
          text="成果物のタイトル"
          subText="20字以内で入力してください。"
          isRequired
          name="title"
          placeholder="成果物共有アプリケーション『Sayka』"
          hasCharCount
        />
        <InputBlock
          text="成果物の内容"
          subText="100字以内で入力してください。"
          isRequired
          name="description"
          hasCharCount
          feature="textarea"
        />
        <ContentSubTitle text="タグ" />
        <Explanation text="成果物に関するタグを最大5つ付与することができます。" />
        <input
          id="tags"
          type="text"
          onChange={handleInputChange}
          placeholder="タグを入力しカンマを入力すると追加されます。"
          className="w-full rounded border border-gray-600 p-3 text-sm focus:bg-indigo-100 focus:outline-indigo-200"
        />
        {tagError && (
          <div className="mt-1 font-bold text-red-600">{tagError}</div>
        )}
        {tags.length >= 5 && !tagError && (
          <div className="mt-1 font-bold text-red-600">
            5つ以上のタグは禁止です。
          </div>
        )}
        <div className="mt-5 flex flex-wrap ">
          {tags.map((tag, index) => (
            <div key={index} className="mr-2 flex items-center py-1">
              <Tag text={tag} />
              <GrClose
                className="ml-1 cursor-pointer"
                size={13}
                onClick={() => handleTagRemove(index)}
              />
            </div>
          ))}
        </div>

        <ContentSubTitle text="関連URL" />
        <Explanation text="全て任意の項目です。" />
        <InputBlock
          text="GitHubなどのソースコード"
          name="github_url"
          placeholder="https://github.com/~"
          icon={<AiFillGithub size={18} />}
        />
        <InputBlock
          text="Figmaなどのデザイン"
          name="figma_url"
          placeholder="https://www.figma.com/~"
          icon={<PiFigmaLogoDuotone size={18} />}
        />
        <InputBlock
          text="Google SlideやCanvaなどのプレゼンテーション資料"
          name="slide_url"
          placeholder="https://docs.google.com/~"
          icon={<TfiLayoutSlider size={18} />}
        />
        <InputBlock
          text="Qiitaやドキュメント資料"
          name="article_url"
          placeholder="https://qiita.com/~"
          icon={<MdOutlineArticle size={18} />}
        />
        <InputBlock
          text="作成したアプリケーションのURL"
          name="app_url"
          placeholder="https://sayka.vercel.app"
          icon={<TfiWorld size={18} />}
        />
        <div className="flex justify-center pt-5">
          <button
            type="button"
            onClick={methods.handleSubmit(onSubmit)}
            className="w-36 rounded bg-sc px-4 py-2 font-semibold text-white transition duration-300 hover:bg-hover-sc">
            投稿する
          </button>
        </div>
      </form>
    </FormProvider>
  );
};
