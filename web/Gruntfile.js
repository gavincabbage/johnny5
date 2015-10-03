// Gruntfile.js
module.exports = function(grunt) {

  grunt.initConfig({

    jshint: {
      all: ['public/src/js/**/*.js']
    },

    // take all the js files and minify them into app.min.js
    uglify: {
      build: {
        files: {
          'public/dist/js/app.min.js': ['public/src/js/**/*.js', '!public/src/js/**/*.test.js']
        }
      }
    },

    // take the processed style.css file and minify
    cssmin: {
      build: {
        files: {
          'public/dist/css/style.min.css': ['public/src/css/style.css']
        }
      }
    },

    // watch css and js files and process the above tasks
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

    // watch our node server for changes
    nodemon: {
      dev: {
        script: 'server.js',
        options: {
          ignore: ['config/**', 'public/**']
        }
      }
    },

    // run watch and nodemon at the same time
    concurrent: {
      options: {
        logConcurrentOutput: true
      },
      tasks: ['nodemon', 'watch']
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
    }

  });

  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-cssmin');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-nodemon');
  grunt.loadNpmTasks('grunt-concurrent');
  grunt.loadNpmTasks('grunt-karma');

  grunt.registerTask('dist', ['cssmin', 'jshint', 'uglify']);
  grunt.registerTask('serve', ['dist', 'concurrent']);
  grunt.registerTask('test', ['jshint', 'karma']);
  grunt.registerTask('default', ['dist', 'test']);

};
