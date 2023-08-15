"use client";

import { ErrorMessage } from "@hookform/error-message";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, SubmitHandler, FormProvider } from "react-hook-form";
import { z } from "zod";

const schema = z.object({
  title: z.string().nonempty({ message: "name is required" }),
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
    <div>
      <FormProvider {...methods}>
        <div className="mb-5">
          <label htmlFor="title" className="mb-1">
            <span className="">タイトル</span>
          </label>
          <div className="my-1 text-xs text-red-600">
            <ErrorMessage
              errors={methods.formState.errors}
              name="title"
              render={({ message }) => <p>{message}</p>}
            />
          </div>
          <input id="title" type="text" {...methods.register("title")} />
        </div>
        {/* this */}
        <form onSubmit={methods.handleSubmit(onSubmit)}>
          <button type="submit">投稿する</button>
        </form>
      </FormProvider>
    </div>
  );
};
