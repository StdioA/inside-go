var gulp  = require('gulp');
var uglify = require('gulp-uglify');
var uglifycss = require('gulp-uglifycss');
var rename = require('gulp-rename');
var less = require('gulp-less')

gulp.task("compress-js", [], function () {
	return gulp.src('./mblog/static/mblog/js/!(*.min).js')
			.pipe(uglify())
			.pipe(rename({ suffix: ".min" }))
			.pipe(gulp.dest('./mblog/static/mblog/js/'))
});

gulp.task("less", [], function () {
	return gulp.src('./mblog/static/mblog/less/*.less')
			.pipe(less())
			.pipe(gulp.dest('./mblog/static/mblog/css/'))
})

gulp.task("compress-css", ["less"], function () {
	return gulp.src('./mblog/static/mblog/css/!(*.min).css')
			.pipe(uglifycss())
			.pipe(rename({ suffix: ".min" }))
			.pipe(gulp.dest('./mblog/static/mblog/css/'))
});

gulp.task('watch', ["compress-js", "compress-css"], function () {
	var watcher_js = gulp.watch('./mblog/static/mblog/js/*.js', ['compress-js']);
	var watcher_css = gulp.watch('./mblog/static/mblog/less/*.less', ['less', 'compress-css']);
	
	watcher_js.on('change', function (event) {
		console.log('File ' + event.path + ' was ' + event.type + ', running tasks...');
	});
	
	watcher_css.on('change', function (event) {
		console.log('File ' + event.path + ' was ' + event.type + ', running tasks...');
	});
});

gulp.task('compress', ['compress-js', 'compress-css']);

gulp.task('default', ['compress']);
