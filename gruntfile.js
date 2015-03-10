module.exports = function(grunt){

  'use strict';

  // Force use of Unix newlines
  grunt.util.linefeed = '\n';

  grunt.loadNpmTasks('grunt-contrib-less');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-requirejs');

  // Project configuration.
  grunt.initConfig({

    // Metadata.
    pkg: grunt.file.readJSON('package.json'),

    less: {
      compileCore: {
        options: {
          strictMath: true,
          noColor: true,
          outputSourceFiles: true
        },
        files: {
          'app/public/css/<%= pkg.name %>.css': 'app/public/less/bootstrap.less'
        }
      },
      minify: {
        options: {
          cleancss: true,
          report: 'min'
        },
        files: {
          'app/public/css/<%= pkg.name %>.min.css': 'app/public/css/<%= pkg.name %>.css'
        }
      }
    },

    copy: {
      fonts: {
        src: 'app/public/bower/bootstrap/fonts/*',
        dest: 'app/public/fonts',
        expand: true,
        flatten: true,
        filter: 'isFile'
      }
    },

    watch: {
      less: {
        files: '**/*.less',
        tasks: ['less']
      }
    }
  });

  grunt.registerTask('copyfonts',  ['copy:fonts']);
  grunt.registerTask('default', ['less', 'copy']);
};
