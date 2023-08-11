"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect } from "react";
import { useForm, SubmitHandler, FormProvider } from "react-hook-form";
import { z } from "zod";

import { InputBlock } from "./InputBlock";


const schema = z.object({
  title: z.string().nonempty({ message: "Name is required" }),
});

type FormData = {
  title: string;
};

export const SaykaForm: React.FC = () => {
  const methods = useForm<FormData>({
    resolver: zodResolver(schema),
  });

  const onSubmit: SubmitHandler<FormData> = (data) => {
    console.log(data);
  };

  return (
    <div className="pr-40">
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(onSubmit)}>
          <InputBlock
            text="成果物のタイトル"
            subText="成果物のタイトルを入力してください"
            isRequired
            name="title"
            options={{ required: "必須項目です" }}
            type="text"
            placeholder="成果物共有アプリケーション『Sayka』"
          />
          <button type="submit">Submit</button>
        </form>
      </FormProvider>
    </div>
  );
};
