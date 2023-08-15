"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { FC, useEffect, useState } from "react";
import { useForm, SubmitHandler, FormProvider } from "react-hook-form";
import { AiFillGithub } from "react-icons/ai";
import { GrClose } from "react-icons/gr";
import { TfiLayoutSlider } from "react-icons/tfi";
import { TfiWorld } from "react-icons/tfi";
import { z } from "zod";

import { Tag } from "@/ui/Tag";
import { ContentSubTitle, Explanation } from "@/ui/Text";

import { InputBlock } from "./InputBlock";
import { PiFigmaLogoDuotone } from "react-icons/pi";
import { MdOutlineArticle } from "react-icons/md";
import { mock_saykas } from "@/api";

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
  title: string;
};

type TProps = {
  saykaId: number;
};

export const SaykaEditForm: FC<TProps> = ({ saykaId }) => {
  const [tags, setTags] = useState<string[]>([]);
  const [tagError, setTagError] = useState<string | null>(null);
  const methods = useForm<FormData>({
    resolver: zodResolver(schema),
  });

  // TODO
  const sayka = mock_saykas.find((sayka) => sayka.id === saykaId);

  useEffect(() => {
    if (!sayka) return;
    if (!sayka.tags) return;
    const currentTags = sayka.tags.map((t) => t.name);
    setTags(currentTags);
  }, [sayka]);

  if (!sayka) return null;

  const onSubmit: SubmitHandler<FormData> = (data) => {
    const submitData = {
      ...data,
      tags,
    };
    console.log(submitData);
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

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)}>
        <InputBlock
          text="成果物のタイトル"
          subText="20字以内で入力してください。"
          isRequired
          name="title"
          placeholder="成果物共有アプリケーション『Sayka』"
          defaultValue={sayka?.title}
          hasCharCount
        />
        <InputBlock
          text="成果物の内容"
          subText="100字以内で入力してください。"
          isRequired
          name="description"
          hasCharCount
          defaultValue={sayka?.description}
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
            <div key={index} className="flex items-center py-1 mr-2">
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
          defaultValue={sayka?.github_url}
          placeholder="https://github.com/~"
          icon={<AiFillGithub size={18} />}
        />
        <InputBlock
          text="Figmaなどのデザイン"
          name="figma_url"
          defaultValue={sayka?.figma_url}
          placeholder="https://www.figma.com/~"
          icon={<PiFigmaLogoDuotone size={18} />}
        />
        <InputBlock
          text="Google SlideやCanvaなどのプレゼンテーション資料"
          name="slide_url"
          defaultValue={sayka?.slide_url}
          placeholder="https://docs.google.com/~"
          icon={<TfiLayoutSlider size={18} />}
        />
        <InputBlock
          text="Qiitaやドキュメント資料"
          name="article_url"
          defaultValue={sayka?.article_url}
          placeholder="https://qiita.com/~"
          icon={<MdOutlineArticle size={18} />}
        />
        <InputBlock
          text="作成したアプリケーションのURL"
          name="application_url"
          defaultValue={sayka?.application_url}
          placeholder="https://sayka.vercel.app"
          icon={<TfiWorld size={18} />}
        />
        <div className="flex justify-center pt-5">
          <button
            type="submit"
            className="w-36 rounded bg-sc px-4 py-2 font-semibold text-white transition duration-300 hover:bg-hover-sc">
            編集する
          </button>
        </div>
      </form>
    </FormProvider>
  );
};