"use client";

import { ErrorMessage } from "@hookform/error-message";
import { FC, useState } from "react";
import { useFormContext, RegisterOptions } from "react-hook-form";

type TProps = {
  text: string;
  subText?: string;
  isRequired?: boolean;
  type?: string;
  name: string;
  options?: RegisterOptions;
  placeholder?: string;
  defaultValue?: string | number;
  icon?: React.ReactNode;
  feature?: "input" | "textarea";
  hasCharCount?: boolean;
  onChange?: (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => void;
};

export const InputBlock: FC<TProps> = ({
  text,
  subText,
  isRequired = false,
  type = "text",
  name,
  options,
  placeholder,
  defaultValue,
  icon,
  feature = "input",
  hasCharCount = false,
  onChange,
}) => {
  const {
    register,
    formState: { errors },
  } = useFormContext();

  const colorStyle = errors[name]
    ? "bg-red-100 focus:outline-red-200"
    : "border-gray-600 focus:bg-indigo-100 focus:outline-indigo-200";

  const [charCount, setCharCount] = useState(0);

  const handleCharCountChange = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const charNum = event.target.value.length;
    setCharCount(charNum);
  };
  const handleChange = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    handleCharCountChange(event);
    onChange && onChange(event);
  };

  return (
    <div className="mb-5">
      {icon ? (
        <label htmlFor={name} className="mb-1">
          <div className="flex space-x-1">
            {icon}
            <span className="text-sm font-semibold">
              {text}
              {isRequired && <span className="text-red-600">*</span>}
            </span>
          </div>
          {subText && <div className="text-xs text-gray-400">{subText}</div>}
        </label>
      ) : (
        <label htmlFor={name} className="mb-1">
          <span className="text-sm font-semibold">
            {text}
            {isRequired && <span className="text-red-600">*</span>}
          </span>
          {subText && <div className="text-xs text-gray-400">{subText}</div>}
        </label>
      )}
      <div className="my-1 text-xs text-red-600">
        <ErrorMessage
          errors={errors}
          name={name}
          render={({ message }) => <p>{message}</p>}
        />
      </div>
      {feature === "textarea" ? (
        <textarea
          className={`${colorStyle} ${"w-full rounded border p-3 text-sm"}`}
          id={name}
          {...register(name, options)}
          placeholder={placeholder}
          defaultValue={defaultValue}
          onChange={handleChange}
        />
      ) : (
        <input
          className={`${colorStyle} ${"w-full rounded border p-3 text-sm"}`}
          id={name}
          type={type}
          {...register(name, options)}
          placeholder={placeholder}
          defaultValue={defaultValue}
          onChange={handleChange}
        />
      )}
      {hasCharCount && (
        <div className="mt-2 flex justify-end text-xs text-gray-400">
          {charCount} 文字
        </div>
      )}
    </div>
  );
};
