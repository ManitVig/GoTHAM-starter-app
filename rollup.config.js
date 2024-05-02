import resolve from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';
import postcss from 'rollup-plugin-postcss';
import { terser } from 'rollup-plugin-terser';

const production = !process.env.ROLLUP_WATCH;

export default {
	input: 'static/main.js',
	output: {
		file: './public/bundle.js',
	},
	plugins: [
		resolve(), 
		commonjs(),
		postcss(),
		production && terser() // minify, but only in production
	]
};
