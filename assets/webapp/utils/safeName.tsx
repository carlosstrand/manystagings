import slugify from 'slugify';

export default (value: string): string => {
  return slugify(value, { lower: true, strict: true });
};
