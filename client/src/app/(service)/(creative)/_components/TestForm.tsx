"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { useForm, SubmitHandler, FormProvider } from "react-hook-form";
import { GrClose } from "react-icons/gr";
import { z } from "zod";

import { Tag } from "@/ui/Tag";

import { InputBlock } from "./InputBlock";

const schema = z.object({
  title: z.string().nonempty({ message: "name is required" }),
});

type FormData = {
  title: string;
};

export const TestForm: React.FC = () => {
  const [tags, setTags] = useState<string[]>([]);
  const methods = useForm<FormData>({
    resolver: zodResolver(schema),
  });

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
    <div>
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(onSubmit)}>
          <InputBlock
            text="成果物のタイトル"
            isRequired
            name="title"
            options={{ required: "必須項目です" }}
            type="text"
            placeholder="成果物共有アプリケーション『Sayka』"
          />
          <input
            id="tags"
            type="text"
            onChange={handleInputChange}
            placeholder="タグを入力しカンマを入力すると追加されます。"
            className="w-full rounded border border-gray-600 p-3 text-sm focus:bg-indigo-100 focus:outline-indigo-200"
          />
          {tags.length >= 5 && (
            <div className="mt-1 font-bold text-red-600">
              5つ以上のタグは禁止です。
            </div>
          )}
          <div className="mt-5 flex flex-wrap space-x-3">
            {tags.map((tag, index) => (
              <div key={index} className="flex items-center">
                <Tag text={tag} />
                <GrClose
                  className="ml-1 cursor-pointer"
                  size={13}
                  onClick={() => handleTagRemove(index)}
                />
              </div>
            ))}
          </div>

          <div className="flex justify-center pt-5">
            <button
              type="submit"
              className="w-36 rounded bg-sc px-4 py-2 font-semibold text-white transition duration-300 hover:bg-hover-sc">
              投稿する
            </button>
          </div>
        </form>
      </FormProvider>
    </div>
  );
};
