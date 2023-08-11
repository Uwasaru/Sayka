"use client";

import { ErrorMessage } from "@hookform/error-message";
import { FC } from "react";
import { useFormContext, RegisterOptions } from "react-hook-form";

type TProps = {
  text: string;
  subText?: string;
  isRequired?: boolean;
  type: string;
  name: string;
  options?: RegisterOptions;
  placeholder?: string;
  defaultValue?: string | number;
  unit?: string;
  feature?: "input" | "textarea";
};

export const InputBlock: FC<TProps> = ({
  text,
  subText,
  isRequired = false,
  type,
  name,
  options,
  placeholder,
  defaultValue,
  feature = "input",
}) => {
  const {
    register,
    formState: { errors },
  } = useFormContext();

  const colorStyle = errors[name]
    ? "bg-red-100 focus:outline-red-200"
    : "border-gray-600 focus:bg-indigo-100 focus:outline-indigo-200";

  return (
    <div className="mb-5">
      <label htmlFor={name} className="mb-2">
        <span className="text-sm">
          {text}
          {isRequired && <span className="text-red-600">*</span>}
        </span>
        {subText && <div className="text-xs text-gray-400">{subText}</div>}
      </label>
      <div className="my-2 text-xs text-red-600">
        <ErrorMessage
          errors={errors}
          name={name}
          render={({ message }) => <p>{message}</p>}
        />
      </div>
      {feature === "textarea" ? (
        <textarea
          className={`${colorStyle} ${"mr-4 w-full rounded border px-3 py-2 text-sm"}`}
          id={name}
          {...register(name, options)}
          placeholder={placeholder}
          defaultValue={defaultValue}
        />
      ) : (
        <input
          className={`${colorStyle} ${"mr-4 w-full rounded border px-3 py-2 text-sm"}`}
          id={name}
          type={type}
          {...register(name, options)}
          placeholder={placeholder}
          defaultValue={defaultValue}
        />
      )}
    </div>
  );
};
