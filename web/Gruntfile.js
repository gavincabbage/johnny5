module.exports = function(grunt) {

  grunt.initConfig({

    jshint: {
      all: ['public/src/js/**/*.js']
    },

    uglify: {
      build: {
        files: {
          '../dist/web/app.min.js': ['public/src/js/**/*.js', '!public/src/js/**/*.test.js']
        }
      }
    },

    cssmin: {
      build: {
        files: {
          '../dist/web/style.min.css': ['public/src/css/style.css']
        }
      }
    },

    watch: {
      options: {
        livereload: true,
        interval: 2000,
        spawn: false
      },
      css: {
        files: ['public/src/css/**/*.css'],
        tasks: ['cssmin']
      },
      js: {
        files: ['public/src/js/**/*.js'],
        tasks: ['jshint', 'uglify']
      },
      jsTest: {
        files: ['public/src/js/**/*.test.js'],
        tasks: ['karma']
      },
      html: {
        files: ['public/**/*.html']
      }
    },

    karma: {
      unit: {
        options: {
          frameworks: ['jasmine'],
          singleRun: true,
          browsers: ['PhantomJS'],
          files: [
            'public/libs/angular/angular.js',
            'public/libs/angular-route/angular-route.js',
            'public/libs/angular-mocks/angular-mocks.js',
            'public/src/js/**/*.js'
          ],
          reporters: ['progress', 'coverage'],
          coverageReporter: {
            reporters: [
              {type: 'text'}
            ]
          },
          preprocessors: { 'public/src/js/**/!(*test).js': ['coverage'] }
        }
      }
    },

    copy: {
        main: {
            files: [
                {
                    cwd: 'public',
                    src: ['templates/*.html', 'index.html'],
                    dest: '../dist/web/',
                    nonull: true,
                    expand: true
                }
            ]
        }
    }

  });

  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-cssmin');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-karma');

  grunt.registerTask('dist', ['cssmin', 'jshint', 'uglify', 'copy']);
  grunt.registerTask('test', ['jshint', 'karma']);
  grunt.registerTask('default', ['test', 'dist']);

};
